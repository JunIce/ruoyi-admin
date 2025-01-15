import Konva from "konva";
import { EventEmitter } from "eventemitter3";
import { GlobalEvent } from "./events";

import { formatValueToNumber, pxUtils } from "./util";

import { createNodeAction } from "./actions";
import Selection from "./selection";
import { Ruler } from "./ruler";

const DesignColumns = [
  {
    label: "Width",
    prop: "width",
    type: "inputNumber",
  },
  {
    label: "Height",
    prop: "height",
    type: "inputNumber",
  },
  {
    label: "ScaleX",
    prop: "scaleX",
    type: "inputNumber",
  },
  {
    label: "ScaleY",
    prop: "scaleY",
    type: "inputNumber",
  },
  {
    label: "单位",
    prop: "unit",
    type: "select",
    disabled: true,
    options: [
      { label: "毫米", value: "mm" },
      { label: "像素", value: "px" },
    ],
  },
];

export default class Designer extends EventEmitter {
  static columns = DesignColumns;

  constructor(config) {
    super();
    this.el = config.container;
    this.config = config;
    this.container = config.container;

    this.stage = new Konva.Stage({
      container: this.config.container,
      width: this.config.container.clientWidth,
      height: this.config.container.clientHeight,
    });

    this.ruler = new Ruler(this);
    this.resizeEnabled = false;
    this.UNIT = config.unit || "mm";

    this.layer = new Konva.Layer({
      width: this.config.width,
      height: this.config.height,
      x: 40,
      y: 40,
    });
    this.stage.add(this.layer);

    this.init();
    this.nodes = [];
    this.selection = new Selection(this);
    this.uploader = config.uploadImage;

    this.actions = createNodeAction(this);
    // 全局事件
    this.globalEvent = new GlobalEvent(this);

    this.bgRect = new PageStage({
      designer: this,
      width: this.config.width,
      height: this.config.height,
      unit: this.UNIT,
    });
    this.layer.add(this.bgRect);
  }

  init() {
    console.log("init");
    this.nodes = [];

    this.ruler.setPosition(this.layer.x(), this.layer.y());

    setTimeout(() => {
      this.emit("init", this.bgRect.getStageAttrs());
    }, 50);
  }

  setOutDirection() {
    let direction = new Konva.Rect({
      width: 4,
      height: this.config.height,
      fill: "#b93f3f",
    });

    this.layer.add(direction);
  }

  // 缩放舞台
  scaleStage(e) {
    let step = 0.05 * (e.deltaY > 0 ? -1 : 1);
    let scale = this.stage.scaleX() + step;
    this.stage.scale({ x: scale, y: scale });
  }

  command({ type }) {
    if (type === "delete") {
      this.nodes = this.nodes.filter(
        (node) => node.id !== this.selectedNode.id
      );
      if (this.selectedNode) {
        let id = this.selectedNode.id;
        this.emit("delete:node", id);
        this.selectedNode.destroy();
        this.selection.unselect();
      }
    } else if (type === "clear") {
      this.nodes.forEach((node) => {
        node.destroy();
      });
      this.nodes = [];
    } else if (type === "resetZoom") {
      this.stage.scale({ x: 1, y: 1 });
    } else if (type === "zoomIn") {
      this.stage.scale({
        x: this.stage.scaleX() + 0.05,
        y: this.stage.scaleY() + 0.05,
      });
    } else if (type === "zoomOut") {
      this.stage.scale({
        x: this.stage.scaleX() - 0.05,
        y: this.stage.scaleY() - 0.05,
      });
    } else {
      this.actions(type);
    }
  }

  // 更新节点属性
  updateNodeAttrs(attrs) {
    if (!this.selectedNode) {
      formatValueToNumber(attrs, [
        "width",
        "height",
        "x",
        "y",
        "scaleX",
        "scaleY",
      ]);
      this.bgRect.setSize(
        this.transformUnitToPx(attrs.width),
        this.transformUnitToPx(attrs.height)
      );
      return;
    }
    this.selectedNode.updateAttrs(attrs);
  }

  // 获取舞台节点
  getStageNodes() {
    // let data = JSON.parse(this.stage.toJSON());
    let stageAttrs = this.bgRect.getStageAttrs();

    let stage = {
      className: "Stage",
      attrs: {
        ...stageAttrs,
      },
      children: [],
    };

    this.nodes.forEach((node) => {
      let attrs = node.getNodeAttrs();
      stage.children.push({
        className: attrs.className,
        attrs,
      });
    });

    return stage;
  }

  fromJSON(json) {
    this.UNIT = json.attrs?.unit || "mm";

    this.bgRect.setSize(
      this.transformUnitToPx(Number(json.attrs?.width)),
      this.transformUnitToPx(Number(json.attrs?.height))
    );

    json.children.forEach((child) => {
      let type = child.className;

      let childAttrs = {
        ...child,
        attrs: child.attrs,
      };

      if (type === "Barcode") {
        childAttrs.attrs.image = window.icons.barcode;
      } else if (type === "QrCode") {
        childAttrs.attrs.image = window.icons.qrcode;
      }

      this.actions(type, childAttrs);
    });

    this.selection.clearSelection();
  }

  get selectedNode() {
    return this.selection.activeNode;
  }

  getSelection() {
    return this.selection.activeNode;
  }

  // 单位转换
  transformUnit(value = 0, unit = this.UNIT) {
    if (unit === "mm") {
      return Number(pxUtils.pxToMm(Number(value))).toFixed(2);
    }
    return Number(value);
  }
  // 单位转换为px
  transformUnitToPx(value = 0, unit = this.UNIT) {
    if (unit === "mm") {
      return pxUtils.mmToPx(Number(value));
    }
    return Number(value);
  }

  destroy() {
    this.nodes.forEach((node) => {
      node.destroy();
    });
    this.nodes = [];
    this.layer.destroyChildren();
    this.stage.destroy();
  }

  getStageAttrs() {
    return this.bgRect.getStageAttrs();
  }
}

// 舞台画布
class PageStage extends Konva.Group {
  constructor(config) {
    super({
      x: 0,
      y: 0,
    });
    this.stage = config.stage;
    this.designer = config.designer;
    this.config = this.designer.config;
    this.trVisible = false;
    this.resizeEnabled = false;
    this.init();
  }

  init() {
    this.bgRect = new Konva.Rect({
      x: 0,
      y: 0,
      width: this.designer.transformUnitToPx(this.config.width),
      height: this.designer.transformUnitToPx(this.config.height),
      fill: "#fff",
      name: "Stage",
      className: "Stage",
      shadowBlur: 20,
      shadowColor: "#ccc",
      draggable: false,
      unit: this.designer.UNIT,
    });
    this.add(this.bgRect);

    this.tr = new Konva.Transformer({
      node: this.bgRect, // 指定要应用 Transformer 的节点
      enabledAnchors: [
        "top-left",
        "top-right",
        "bottom-left",
        "bottom-right",
        "middle-left",
        "middle-right",
        "top-center",
        "bottom-center",
      ],
      visible: this.resizeEnabled,
      rotateEnabled: false,
    });
    this.add(this.tr);

    this.bgRect.on("transform", () => {
      // 获取当前变换后的尺寸
      const scaleX = this.bgRect.scaleX();
      const scaleY = this.bgRect.scaleY();

      this.bgRect.setAttrs({
        x: 0,
        y: 0,
        scaleX: 1,
        scaleY: 1,
        width: this.bgRect.width() * scaleX,
        height: this.bgRect.height() * scaleY,
      });

      this.designer.emit("update:node", this.bgRect, this.getStageAttrs());
    });
  }

  // 获取舞台属性
  getStageAttrs() {
    let attrs = this.bgRect.getAttrs() || {};
    return {
      ...attrs,
      width: this.designer.transformUnit(Number(attrs.width)?.toFixed(2)),
      height: this.designer.transformUnit(Number(attrs.height)?.toFixed(2)),
      scaleX: Number(attrs.scaleX || 1)?.toFixed(2) || 1,
      scaleY: Number(attrs.scaleY || 1)?.toFixed(2) || 1,
    };
  }

  setSize(width, height) {
    this.bgRect.setAttrs({
      width,
      height,
    });
  }

  toggleTr(visible) {
    this.trVisible = visible;
    this.tr.visible(visible);
  }
}
