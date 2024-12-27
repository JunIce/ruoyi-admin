// Vue2 组件, 获取未读消息
import { listNotice } from "@/api/system/notice";
import Vue from "vue";
let noticeInstanceMap = new Map();

export default {
  mounted() {
    this.fetchNotice();
  },
  methods: {
    async fetchNotice() {
      const h = this.$createElement;
      // 获取未读消息 noticeType: 101 类型
      const { rows } = await listNotice({
        pageSize: 10,
        noticeType: 101,
      });
      let unreadNotices = rows.filter((row) => row.status == 0);
      let taskrun = () => {
        this.timer = setTimeout(this.fetchNotice.bind(this), 2000);
      };

      // 消息提示
      let noticeTask = (row) => {
        // 如果已经存在实例，则不创建新的实例
        if (noticeInstanceMap.has(row.id)) {
          return;
        }
        let instance;
        let renderComp = (msg) => {
          return Vue.extend({
            render() {
              return h("div", [
                h("p", { style: "color: #333" }, msg),
                h("div", { style: "text-align: right" }, [
                  h(
                    "el-button",
                    {
                      props: { type: "primary", size: "small" },
                      style: {
                        textAlign: "right",
                      },
                      on: {
                        click: () => {
                          updateNotice({
                            ...row,
                            status: 1,
                          });
                          instance.close();
                          noticeInstanceMap.delete(row.id);
                        },
                      },
                    },
                    "确定"
                  ),
                ]),
              ]);
            },
          });
        };

        instance = this.$notify({
          title: "提示",
          message: h(renderComp(row.noticeContent)),
          duration: 0,
          customClass: "custom-notifyNotification",
        });
        noticeInstanceMap.set(row.id, instance);
      };

      // 未读消息提示
      unreadNotices.forEach(noticeTask);
      // 启动任务
      taskrun();
    },
  },
};
