import { EventEmitter } from 'eventemitter3';

export class DH_RPC extends EventEmitter {
  constructor(ip, port) {
    super();

    this.ip = ip;
    this.port = port;
    this.target = ip + ':' + port;
    this.player = null;
    setIP(this.target);
    this.status = 0;
    this.loading = false;
  }
  Login({ username, password }) {
    this.username = username;
    this.password = password;
    /**
     * RPC.login 登录
     * @param {string} $user.value 用户名
     * @param {string} $password.value 密码
     * @param {boolean} false 是否httpOnly,默认false
     * @returns {Promise}
     */
    return RPC.login(username, password, false)
      .then((res) => {
        console.info('登录成功', this.target);
        setCookie('DWebClientSessionID', '', -1);
        setCookie('DhWebClientSessionID', '', -1);
        /**
         * RPC.keepAlive 保活
         */
        // debugger;
        RPC.keepAlive(300, 60000, _getSession(), this.target);
        const browser = BrowserType();
        if (browser.includes('ie')) {
          window.onunload = () => {
            ajax({
              url: 'global.logout',
            });
          };
        } else if (browser.includes('chrome')) {
          const params = {
            method: 'global.logout',
            params: null,
            id: 10000,
            session: _getSession(),
          };
          pubsub.subscribe('onbeforeunload', () => {
            navigator.sendBeacon('/RPC2', JSON.stringify(params));
          });
        } else {
          pubsub.subscribe('onbeforeunload', () => {
            ajax({
              url: 'global.logout',
            });
          });
        }
        return res;
        // $loginState.style.color = 'green';
        // $loginState.innerHTML = '已登录';
        // setLoginState(true);
        // afterLogin();
      })
      .catch((err) => {
        console.log(err);
        loginError(err);
      });
  }

  // 获取所有摄像头列表
  getAllCameraList() {
    return RPC.LogicDeviceManager.getCameraAll().then(function (params) {
      let channelList = params.camera.filter((item) => item.Enable === true);
      //预览，在线通道列表
      let channelArr = channelList.map((item) => {
        let _name;
        item.DeviceInfo.VideoInputs.map((value) => {
          if (value && value.Enable) {
            _name = value.Name;
          }
        });
        return item.UniqueChannel + ';' + _name;
      });
      // 回放，全部通道列表
      let allArr = params.camera
        .map((item) => {
          let _name;
          if (item.DeviceInfo && item.DeviceInfo.VideoInputs) {
            item.DeviceInfo.VideoInputs.map((value) => {
              _name = value && value.Name;
            });
            return (
              item.UniqueChannel + ';D' + (item.UniqueChannel + 1) + ' ' + _name
            );
          }
        })
        .filter((x) => x);

      return {
        ...params,
        channelArr,
        allArr,
      };
    });
  }

  // 码流列表
  getProductDefinition() {
    return RPC.MagicBox.getProductDefinition('MaxExtraStream').then(function (
      params,
    ) {
      let maxExtra = params.definition;
      let list = [
        {
          value: 0,
          label: '主码流',
        },
      ];

      if (maxExtra > 1) {
        for (let i = 1; i <= maxExtra; i++) {
          list.push({ label: '辅码流' + i, value: i });
        }
      } else {
        list.push({ label: '辅码流', value: 1 });
      }

      return list;
    });
  }

  // 获取时间段内视频列表
  async searchPlaybackList({ startTime, endTime, channel }) {
    let allRecords = [];
    let recordNums = 0;
    // let playChannel = $('#h5_playback_channel').value - 0;
    const getMediaFile = (params) => {
      return new Promise((resolve, reject) => {
        /**
         * RPC.MediaFileFind.instance 创建媒体文件查找实例
         * @returns {Promise}
         */
        RPC.MediaFileFind.instance()
          .then((json) => {
            let queryId = json.result;
            /**
             * RPC.MediaFileFind.findFile 设置查找条件，并判断是否存在文件
             * @param {number} queryId 实例id
             * @param {object} params condition参数
             * @returns {Promise}
             */
            RPC.MediaFileFind.findFile(queryId, params)
              .then(() => {
                findNextFile(queryId)
                  .then((data) => {
                    console.log(11111, data);
                    resolve(data);
                  })
                  .catch((err) => {
                    reject(err);
                  });
              })
              .catch((err) => {
                reject(err);
              });
          })
          .catch((err) => {
            reject(err);
          });
      });
    };
    const findNextFile = (queryId) => {
      return new Promise((resolve, reject) => {
        /**
         * RPC.MediaFileFind.findNextFile 在指定条件基础上查询文件信息
         * @param {number} queryId 实例
         * @param {object} 需要查找的数目
         * @returns {Promise}
         */
        RPC.MediaFileFind.findNextFile(queryId, { count: 100 })
          .then((data) => {
            if (Number.isInteger(data.found)) {
              recordNums = recordNums + data.found;
              allRecords = allRecords.concat([...data.infos]);
              console.log('allRecords:', allRecords);
              if (data.found === 100) {
                findNextFile(queryId)
                  .then(() => {
                    resolve(true);
                  })
                  .catch((err) => {
                    reject(err);
                  });
              } else {
                stopFind(queryId);
                resolve([...allRecords]);
              }
            } else {
              stopFind(queryId);
              resolve([]);
            }
          })
          .catch((err) => {
            reject(err);
            stopFind(queryId);
          });
      });
    };
    const stopFind = (object) => {
      return new Promise((resolve, reject) => {
        /**
         * PC.MediaFileFind.close 结束查询
         * @param {number} object 媒体文件查找实例ID
         * @returns {Promise}
         */
        RPC.MediaFileFind.close(object)
          .then(() => {
            /**
             * PC.MediaFileFind.destroy 销毁媒体文件查找实例
             * @param {number} object 媒体文件查找实例ID
             */
            RPC.MediaFileFind.destroy(object);
            resolve(true);
          })
          .catch(() => {
            reject();
          })
          .finally(() => {});
      });
    };
    // const updateInfos = (infos) => {
    //     let table =  document.querySelector('#h5_table tbody');
    //     table.innerHTML = '';
    //     for(let i = 0; i < infos.length; i++) {
    //         let time = infos[i].StartTime + ' - ' + infos[i].EndTime;  //<input type="button" class="h5-button" btn-for="onGoTime" value="GO!">
    //         let size = Math.round(infos[i].Length / 1024);
    //         let newRow = table.insertRow(-1);
    //         newRow.innerHTML = `<td><input type="checkbox" id="h5_check_${i}"></td><td>${i+1}</td><td>${time}</td><td>${size}</td><td><span id="h5_curTime_${i}">--</span><span>/</span><span id="h5_totalTime_${i}">--</span><input type="text" id="h5_goTime_${i}" style="width: 50px;"><input type="button" class="h5-button" id="h5_button_go_${i}" value="GO!"></td>`;
    //     }
    //     document.querySelectorAll('[id^=h5_button_go_]').forEach(item => {
    //         item.addEventListener('click', function(event) {
    //             let id = item.getAttribute('id').split('_')[3] - 0;
    //             onGoTime(id);
    //         });
    //     });
    //     document.querySelectorAll('[id^=h5_check_]').forEach(function(item) {
    //         item.addEventListener('click', function(event) {
    //             event.stopPropagation();
    //             if(event.target.checked) {
    //                 //渲染裁剪时间
    //                 let _index = event.target.getAttribute('id').split('_')[2] - 0;
    //                 let startTime = recordArr[_index + lINENUMBER*(curPage-1)].StartTime.split(' ').join('T');
    //                 let endTime = recordArr[_index + lINENUMBER*(curPage-1)].EndTime.split(' ').join('T');
    //                 if(startTime.split(':')[2] === '00') {
    //                     startTime = startTime.substr(0, startTime.length - 3);
    //                 }
    //                 if(endTime.split(':')[2] === '00') {
    //                     endTime = endTime.substr(0, endTime.length - 3);
    //                 }
    //                 $('#h5_cutStartTime').value = startTime;
    //                 $('#h5_cutEndTime').value = endTime;
    //                 document.querySelector('[btn-for=onStartCut]').setAttribute('cutIndex', _index);
    //             }
    //         });
    //     });
    //     document.querySelectorAll('#h5_table tbody tr').forEach(function(item) {
    //         item.addEventListener('dblclick', function(event) {
    //             event.stopPropagation();
    //             if(event.target.nodeName === 'TD') {
    //                 event.target.style.color = 'blue';
    //                 let dom = event.target.parentNode.childNodes[1];
    //                 let value = dom.innerText - 1;
    //                 let url = recordArr[value].FilePath;
    //                 onStopPreview();
    //                 onPreview(true, url, value);
    //             }
    //         });
    //     });
    // }
    // const updatePages = () => {
    //     totalPage = Math.ceil(recordNums/lINENUMBER);
    //     $('#h5_curPage').innerText = curPage;
    //     $('#h5_totalPage').innerText = totalPage;
    // }
    let tmpDir = [];
    try {
      /**
       * RPC.getDeviceAllInfo 获取存储信息
       * @param {string} 'getDeviceAllInfo' 方法名
       * @return {Promise}
       */
      tmpDir = await RPC.getDeviceAllInfo('getDeviceAllInfo');
    } catch (e) {
      console.log(e);
    }
    let dirs = null;
    if (tmpDir.info && tmpDir.info.length > 1) {
      dirs = 'All';
    } else {
      //dirs = tmpDir.info?[0]?.Detail?[0]?.Path ?? '/mnt/sd';
      dirs =
        (tmpDir.info &&
          tmpDir.info[0] &&
          tmpDir.info[0].Detail &&
          tmpDir.info[0].Detail[0] &&
          tmpDir.info[0].Detail[0].Path) ||
        '/mnt/sd';
    }

    // let startTime = $('#h5_startTime').value.replace('T', ' ');
    // let endTime = $('#h5_endTime').value.replace('T', ' ');
    // if(startTime.split(' ')[1].split(':').length < 3) {
    //     startTime = startTime + ':00';
    // }
    // if(endTime.split(' ')[1].split(':').length < 3) {
    //     endTime = endTime + ':00';
    // }
    let params = {
      condition: {
        Channel: channel,
        Dirs: [dirs],
        StartTime: startTime,
        EndTime: endTime,
        Flags: null,
        Events: ['*'],
        Types: ['dav'],
      },
    };
    return getMediaFile(params).catch((err) => {
      if (err && err.error && err.error.code === 285409409) {
        alert('回放功能需要确保SD卡经过设备认证');
      } else {
        alert('无数据');
      }
    });
  }

  previewFileUrl(
    isPlayback,
    url,
    playbackIndex,
    { videoRef, canvasRef, channel },
  ) {
    // if (
    //   playerInstance[WndIndex] &&
    //   onlineChannel.indexOf(channel) > -1 &&
    //   !isChangeStream
    // ) {
    //   alert('通道' + (channel + 1) + '已存在！');
    //   return;
    // }
    // onStopPreview();
    // var player = null;
    // if (!isLogin) {
    //   alert('请先登录再预览！');
    //   return;
    // }

    if (this.player) {
      this.player.stop();
      this.player = null;
    }

    this.loading = true;

    let curChannel = channel + 1; //无插件通道号从1开始
    let stream = 2;
    let firstTime = 0;
    let ip = this.ip;
    let port = this.port - 0;
    let username = this.username;
    let password = this.password;
    let options = {
      wsURL: 'ws://' + ip + ':' + port + '/rtspoverwebsocket',
      rtspURL: !isPlayback
        ? 'rtsp://' +
          ip +
          ':' +
          port +
          '/cam/realmonitor?channel=' +
          curChannel +
          '&subtype=' +
          stream +
          '&proto=Private3'
        : 'rtsp://' + ip + ':' + port + '/' + url,
      username: username,
      password: password,
      lessRateCanvas: true,
      playback: isPlayback,
      isPrivateProtocol: false,
      realm: RPC.realm, //设备登录返回的realm
      playbackIndex: playbackIndex,
      h265AccelerationEnabled: true, // 硬解码是否开启，默认不开启
    };
    let player = new PlayerControl(options);
    player.on('MSEResolutionChanged', function (e) {
      console.log(e);
    });
    player.on('PlayStart', (e) => {
      console.log(e);
      this.status = 1;
      this.loading = false;
      //   $videoLoading.style.display = 'none';
      //   let curWndType =
      //     document.querySelector('[sel-for=onChangeWdnNum]').value - 0;
      //   if (!player.isPlayback) {
      //     onlineChannel.push(channel);
      //     updateChannelList();
      //     // if(curWndType !== 1) {
      //     //     clickNextWnd();
      //     // }
      //   }
    });
    player.on('DecodeStart', function (e) {
      console.log(e);
      if (e.decodeMode === 'video') {
        videoRef.style.display = 'block';
        canvasRef.style.display = 'none';
      } else {
        videoRef.style.display = 'none';
        canvasRef.style.display = 'block';
      }
      //   canvasSon = new PluginCanvasES6();

      //   canvasSon.init($canvas_ivs, function (data) {
      //     rebackActivateLocalEnlarging(data);
      //   });
      //   canvasSon.addChangeShapeEvent();
      //   playerInstance[WndIndex] = player;
      //   ivsInstance[WndIndex] = canvasSon;
    });
    player.on('UpdateCanvas', (e) => {
      //   if (player.isPlayback) {
      this.emit('updateTimeStamp', e);

      //     let playbackIndex = player.playbackIndex;
      //     if (firstTime === 0) {
      //       firstTime = e.timestamp;
      //     }
      //     //const _left = e.timestamp - new Date(recordArr[playbackIndex].StartTime).getTime()/1000;
      //     $('#h5_curTime_' + (playbackIndex % lINENUMBER)).innerText =
      //       e.timestamp - firstTime;
      //   }
    });
    player.on('GetTotalTime', (e) => {
      console.log('total time: ', e);
      //   let playbackIndex = player.playbackIndex % lINENUMBER;
      //   $('#h5_totalTime_' + playbackIndex).innerText = e;
    });
    player.on('UpdateTimeStamp', (e) => {
      //   console.log('timestamp: ', e);
      //   let playbackIndex = player.playbackIndex % lINENUMBER;
      //   $('#h5_totalTime_' + playbackIndex).innerText = e;
    });
    player.on('GetFrameRate', function (e) {
      console.log('GetFrameRate: ' + e);
    });
    player.on('FrameTypeChange', function (e) {
      console.log('FrameTypeChange: ' + e);
    });
    player.on('Error', function (e) {
      //console.log('Error: ' + JSON.stringify(e))
    });
    player.on('IvsDraw', function (e) {
      //console.log('IvsDraw: ' + JSON.stringify(e))
    });
    player.on('WorkerReady', function () {
      player.connect();
    });
    player.on('Error', () => {
      // player.connect();
      this.loading = false;
    });
    player.init(canvasRef, videoRef);
    this.player = player;
  }

  // 播放到指定时间
  playAtTime(curTime) {
    this.status = 1;
    this.player.playByTime(curTime);
  }

  play() {
    this.status = 1;
    this.player.play();
  }

  pause() {
    this.status = 0;
    this.player.pause();
  }

  destroy() {
    this.player?.close();
    this.player = null;
    RPC.logout();
  }
}
