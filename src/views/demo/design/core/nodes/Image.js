import Konva from "konva";
import { AbstractNode } from "./AbstractNode";
import { loadImage } from "../util";
import Invalid from '../../assets/invalid.png';

const imageCache = {};

export class ImageNode extends AbstractNode {
  type = "Image";

  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      width: 30,
      height: 30,
    };
  }

  static create(designer) {
    let text = new Konva.Image({
      ...this.commonAttrs,
      draggable: true,
      resize: true,
      draggable: true,
      image: '',
      className: "Image",
    });

    let node = new ImageNode(text, designer);

    node.draw(designer.layer);

    return node;
  }

  updateAttrs(attrs) {
    super.updateAttrs(attrs);
    if (attrs.imageUrl) {
      if (imageCache[attrs.imageUrl]) {
        this.node.image(imageCache[attrs.imageUrl]);
      } else {
        loadImage(process.env.VUE_APP_BASE_API + attrs.imageUrl).then(
          (image) => {
            this.node.image(image);
            imageCache[attrs.imageUrl] = image;
          }
        ).catch((e) => {
          console.log(e);
          const invalidImg = new Image();
          invalidImg.src = Invalid;
          invalidImg.onload = () => {
            this.node.image(invalidImg);
          };
        });
      }
    }
  }

  get columns() {
    return [
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
        label: "X",
        prop: "x",
        type: "inputNumber",
      },
      {
        label: "Y",
        prop: "y",
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
        label: "图片",
        prop: "image",
        type: "image",
      },
      {
        label: "绑定属性",
        prop: "bindProp",
        type: "input",
      },
    ];
  }
}
