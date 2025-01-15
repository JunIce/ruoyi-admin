import Konva from "konva";
import { mmToPx } from "./util";

export class Ruler {
  constructor(parent) {
    this.parent = parent;

    this.layer = new Konva.Layer({
      listening: false,
      draggable: false,
      name: "ruler",
      x: 0,
      y: 0,
    });
    this.parent.stage.add(this.layer);
    this.step = Number(mmToPx(1).toFixed(2));

    this.group = new Konva.Group({
      x: 0,
      y: 0,
    });
    this.layer.add(this.group);

    this.createLine();
    this.createXAxisRuler();
    this.createYAxisRuler();
  }

  createLine() {
    let width = this.parent.stage.width();
    let height = this.parent.stage.height();
    let node = new Konva.Shape({
      x: 0,
      y: 0,
      width: width,
      height: height,
      stroke: "#ccc",
      strokeWidth: 0.5,
      sceneFunc: (context, shape) => {
        const step = 20;
        const yCount = height / this.parent.stage.scaleY() / step;
        const xCount = width / this.parent.stage.scaleX() / step;

        context.lineWidth = shape.strokeWidth();
        context.strokeStyle = shape.stroke();
        context.setLineDash([5, 2]);

        for (let i = 1; i < yCount; i++) {
          context.beginPath();
          context.moveTo(0, i * step);
          context.lineTo(width / this.parent.stage.scaleX(), i * step);
          context.stroke();
          context.closePath();
        }

        for (let i = 1; i < xCount; i++) {
          context.beginPath();
          context.moveTo(i * step, 0);
          context.lineTo(i * step, height / this.parent.stage.scaleY());
          context.stroke();
          context.closePath();
        }
      },
    });
    this.group.add(node);
  }

  createXAxisRuler() {
    let step = this.step;
    let xAxis = new Konva.Shape({
      x: 0,
      y: 0,
      lineWidth: 0.5,
      stroke: "#000",
      sceneFunc: (context, shape) => {
        let width = this.parent.stage.width() / this.parent.stage.scaleX();
        let height = 12;

        context.lineWidth = 1;
        context.strokeStyle = "#777";
        context.font = "11px Arial";
        context.fillStyle = "#333";
        context.textBaseline = "middle";

        let start = 0;
        let end = width;
        let count = (end - start) / step;

        for (let i = 0; i < count; i++) {
          context.beginPath();
          context.moveTo(i * step, 0);

          if (i % 10 === 0) {
            context.lineTo(i * step, height);
          } else {
            context.lineTo(i * step, 8);
          }

          context.stroke();
          context.closePath();

          if (i % 10 === 0) {
            let textWidth = context.measureText(i).width;
            context.fillText(i, i * step - textWidth / 2, 25);
          }
        }
      },
    });
    this.xAxis = xAxis;
    this.group.add(xAxis);
  }

  createYAxisRuler() {
    let step = this.step;
    let yAxis = new Konva.Shape({
      x: 0,
      y: 0,
      lineWidth: 0.5,
      stroke: "#000",
      sceneFunc: (context, shape) => {
        let width = this.parent.stage.height() / this.parent.stage.scaleY();
        let length = 12;

        context.lineWidth = 1;
        context.strokeStyle = "#777";
        context.font = "11px Arial";
        context.fillStyle = "#333";
        context.textBaseline = "middle";

        let start = 0;
        let end = width;
        let count = (end - start) / step;

        for (let i = 0; i < count; i++) {
          context.beginPath();
          context.moveTo(0, i * step);

          if (i % 10 === 0) {
            context.lineTo(length, i * step);
          } else {
            context.lineTo(8, i * step);
          }

          context.stroke();
          context.closePath();

          if (i % 10 === 0) {
            context.fillText(i, length + 5, i * step);
          }
        }
      },
    });
    this.group.add(yAxis);
    this.yAxis = yAxis;
  }

  setPosition(x, y) {
    this.xAxis.x(x);
    this.yAxis.y(y);
  }

  rerender() {
    this.layer.draw();
  }
}
