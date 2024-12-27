import { h } from "vue";
import { ElButton, ElNotification } from "element-plus";
let noticeInstanceMap = new Map();

let sleep = (ms) => {
  return new Promise((resolve) => setTimeout(resolve, ms));
};

export function useNotice() {
  const fetchNotice = async (rows) => {
    let unreadNotices = rows.filter((row) => row.status == 0);

    unreadNotices.forEach(async (row) => {
      if (noticeInstanceMap.has(row.id)) {
        return;
      }
      await sleep(50);
      let instance;
      let renderNotification = (msg) => {
        instance = ElNotification({
          title: "消息通知",
          duration: 0,
          showClose: false,
          message: () =>
            h("div", null, [
              h("div", null, msg),
              h("div", { style: "text-align: right" }, [
                h(
                  ElButton,
                  {
                    type: "primary",
                    size: "small",
                    onClick: () => {
                      instance?.close();
                      noticeInstanceMap.delete(row.id);
                    },
                  },
                  "关闭"
                ),
              ]),
            ]),
          type: "warning",
        });

        noticeInstanceMap.set(row.id, instance);
      };

      renderNotification(row.content);
    });
  };

  return {
    noticeInstanceMap,
    fetchNotice,
  };
}
