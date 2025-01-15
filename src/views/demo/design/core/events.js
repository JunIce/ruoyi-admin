export class GlobalEvent {
  constructor(designer) {
    this.designer = designer;

    this.destroys = [];

    this.destroys.push(this.init());
    this.destroys.push(this.keyboard());
    this.destroys.push(this.mouse());
    this.destroys.push(this.stageEvent());
  }

  init() {
    this.designer.on("init", (data) => {});

    this.designer.on("select:node", (node, data) => {
      this.designer.selection.select(node);
      if (!this.designer.nodes.find((item) => item.id === node.id)) {
        this.designer.nodes.push(node);
      }
    });

    this.designer.on("delete:node", (id) => {
      console.log("delete:node", id);
    });

    this.designer.on("update:node", (id, data) => {
      // console.log('update:node', id, data);
    });

    this.designer.on("resize:stage", (bool) => {
      this.designer.bgRect.toggleTr(bool);
    });

    return () => {
      this.designer.off("init");
      this.designer.off("selected");
      this.designer.off("delete:node");
      this.designer.off("update:node");
    };
  }

  keyboard() {
    let keyboardEvent = (e) => {
      if (
        e.target.tagName === "INPUT" ||
        e.target.tagName === "TEXTAREA" ||
        e.target.tagName === "SELECT"
      ) {
        return;
      }
      if (e.key === "Escape") {
        this.designer.command({ type: "cancel" });
      } else if (e.key === "Enter") {
        this.designer.command({ type: "confirm" });
      } else if (e.key === "Delete" || e.key === "Backspace") {
        this.designer.command({ type: "delete" });
      }
      console.log("keyboardEvent", e);
    };
    window.addEventListener("keyup", keyboardEvent);

    return () => {
      window.removeEventListener("keyup", keyboardEvent);
    };
  }

  mouse() {
    let mouseEvent = (e) => {
      let parent = this.designer.el.parentElement;
      if (e.target === parent) {
        this.designer.selection.clearSelection();
      }
    };
    window.addEventListener("mouseup", mouseEvent, false);

    let scrollEvent = (e) => {
      e.preventDefault();
      this.designer.scaleStage(e);
    };

    this.designer.el.addEventListener("wheel", scrollEvent, false);

    return () => {
      window.removeEventListener("mouseup", mouseEvent);
      this.designer.el.removeEventListener("wheel", scrollEvent);
    };
  }

  stageEvent() {
    let stageEvent = (e) => {
      if (e.target?.name() === "Stage") {
        this.designer.selection.clearSelection();
        this.designer.emit("resize:stage", true);
      } else {
        this.designer.emit("resize:stage", false);
      }
    };
    this.designer.stage.on("click", stageEvent);

    return () => {
      this.designer.stage.off("click", stageEvent);
    };
  }

  destroy() {
    this.destroys.forEach((destroy) => destroy?.());
  }
}
