<template>
  <div class="nvr-content">
    <el-row style="height: 100%" :gutter="30">
      <el-col :span="12">
        <el-form :model="form" label-width="80px" inline>
          <el-form-item prop="timeRange" label="时间范围">
            <el-date-picker v-model="form.timeRange" type="datetimerange" value-format="YYYY-MM-DD HH:mm:ss"
              start-placeholder="请选择开始时间" end-placeholder="请选择结束时间"></el-date-picker>
          </el-form-item>
          <el-form-item prop="channel" label="通道">
            <el-select v-model="form.channel" style="width: 220px">
              <el-option v-for="item in cameraList" :key="item.value" :label="item.label"
                :value="item.value"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="getList">查 询</el-button>
          </el-form-item>
        </el-form>

        <div>
          <el-table :data="list" border :max-height="600" :height="600">
            <el-table-column label="索引" type="index" width="60"></el-table-column>
            <el-table-column label="时间" prop="time">
              <template #default="scope">
                {{ scope.row.StartTime }} - {{ scope.row.EndTime }}
              </template>
            </el-table-column>
            <el-table-column label="文件大小(M)" type="index" width="120">
              <template #default="scope">
                {{ Math.round(scope.row.Length / (1024 * 1024)) }}
              </template>
            </el-table-column>
            <el-table-column label="操作" width="100">
              <template #default="scope">
                <el-button type="primary" link @click="playVideo(scope.row)">
                  播放
                </el-button>
                <!-- <el-button type="danger" @click="deleteVideo(scope.row)">
                  删除
                </el-button> -->
              </template>
            </el-table-column>
          </el-table>
        </div>
      </el-col>
      <el-col :span="12" style="height: 100%">
        <div class="h5-play-wrap" v-loading="dhInstance.loading">
          <video class="player-video" ref="videoRef" />
          <canvas class="player-canvas" ref="canvasRef"></canvas>

          <div class="slider">
            <el-icon size="40" @click="handleVideoPlay">
              <VideoPlay v-if="dhInstance.status == 0" />
              <VideoPause v-if="dhInstance.status == 1" />
            </el-icon>
            <el-slider v-model="playCurrentTime" :min="0" :max="maxSecs" :format-tooltip="formatTooltip" :step="5"
              :marks="marks" @change="handlePlayTimeChange" />
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { dayjs } from 'element-plus';
import { DH_RPC } from './rpc';
import { throttle } from 'lodash-es';

export default {
  data() {
    return {
      dhInstance: new DH_RPC(
        import.meta.env.VITE_RPC_IP,
        import.meta.env.VITE_RPC_PORT,
      ),
      list: [], // Create an array of 10 v-cards
      cameraList: [], // 通道列表
      IP: import.meta.env.VITE_RPC_IP,
      username: import.meta.env.VITE_RPC_USERNAME,
      password: import.meta.env.VITE_RPC_PASSWORD,

      maxSecs: 100,
      playCurrentTime: 0,
      form: {},
      changeTime: false,
    };
  },
  computed: {
    marks() {
      return {
        0: { label: '0', value: 0 },
        ...this.generateMarks(),
      };
    },
  },
  mounted() {
    this.init(); // Fetch device list when the component is mounted
    this.dhInstance.on(
      'updateTimeStamp',
      throttle((e) => {
        this.playCurrentTime = Math.floor(
          (dayjs(e.timestamp * 1000) - this.startSecs) / 1000,
        );
      }, 500),
    );
  },
  beforeDestroy() {
    this.dhInstance.off('updateTimeStamp');
    this.dhInstance.destroy();
  },
  methods: {
    async init() {
      this.dhInstance
        .Login({
          username: this.username,
          password: this.password,
          ip: this.IP,
        })
        .then((res) => {
          return this.dhInstance.getProductDefinition();
        })
        .then((res) => {
          return this.dhInstance.getAllCameraList();
        })
        .then((res) => {
          console.log(res);
          this.cameraList = res.allArr.map((text) => {
            return {
              label: text.split(';')[1],
              value: text.split(';')[0],
            };
          });
        });
    },

    // 查询列表
    getList() {
      if (!this.form.timeRange || this.form.timeRange.length == 0) {
        this.$message.warning('请选择开始时间和结束时间');
        return;
      }

      if (!this.form.channel) {
        this.$message.warning('请选择通道');
        return;
      }

      this.dhInstance
        .searchPlaybackList({
          startTime: this.form.timeRange[0],
          endTime: this.form.timeRange[1],
          channel: Number(this.form.channel),
        })
        .then((list) => {
          console.log(list);
          this.list = list;
        });
    },

    formatSeconds(seconds) {
      const hours = Math.floor(seconds / 3600);
      const minutes = Math.floor((seconds % 3600) / 60);
      const remainingSeconds = seconds % 60;

      // 格式化输出，保证两位数字
      return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(
        2,
        '0',
      )}:${String(remainingSeconds).padStart(2, '0')}`;
    },

    playVideo(row) {
      let startSecs = dayjs(row.StartTime);
      this.startSecs = startSecs;
      let endSecs = dayjs(row.EndTime);
      // console.log(startSecs, endSecs);
      this.maxSecs = Math.floor((endSecs - startSecs) / 1000);

      this.dhInstance.previewFileUrl(true, row.FilePath, row.$index, {
        channel: this.form.channel,
        canvasRef: this.$refs.canvasRef,
        videoRef: this.$refs.videoRef,
      });
    },
    //
    formatTooltip(val) {
      return this.formatSeconds(val);
    },
    // 生成mark
    generateMarks() {
      const marks = {};
      const interval = Math.floor(this.maxSecs / 4);

      let i = 0;

      while (i <= this.maxSecs) {
        marks[i] = {
          label: this.formatSeconds(i),
          value: i,
        };
        i += interval;
      }

      return marks;
    },

    handleVideoPlay() {
      if (!this.dhInstance.player) {
        this.$message.warning('请选择播放视频');
        return;
      }
      if (this.dhInstance.status == 0) {
        this.dhInstance.play();
      } else {
        this.dhInstance.pause();
      }
    },

    handlePlayTimeChange() {
      this.dhInstance.playAtTime(this.playCurrentTime);
    },
  },
};
</script>
<style lang="scss">
.nvr-content {
  width: 100%;
  height: 100%;
  padding: 20px;
  background-color: #fff;
  overflow-y: auto;
  // border: 1px solid #ccc;

  .h5-play-wrap,
  .player-canvas {
    width: 100%;
    height: 100%;
  }

  .player-video {
    display: none;
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0;
    left: 0;
  }

  .h5-play-wrap {
    background-color: #000;
    position: relative;
    width: 100%;
    height: 100%;
  }

  .slider {
    position: absolute;
    bottom: 15px;
    left: 0;
    width: 100%;
    padding-right: 30px;
    padding-left: 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .el-icon {
      color: #fff;
      width: 40px;
      height: 40px;
      cursor: pointer;
    }

    .el-slider {
      margin-left: 25px;
      flex: 1;
    }

    .el-slider__marks-text {
      color: #f4f4f4;
    }
  }
}
</style>
