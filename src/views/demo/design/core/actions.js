import { TextNode } from "./nodes/Text";
import { ImageNode } from "./nodes/Image";
import { LineNode } from "./nodes/Line";
import { BarcodeNode } from "./nodes/Barcode";
import { RectNode } from "./nodes/Rect";
import { CircleNode } from "./nodes/Circle";
import { QrCodeNode } from "./nodes/QrCode";
import { loadImage } from "./util";

export function createAction(designer) {
  return (type) => {
    return {
      type,
      designer,
    };
  };
}

export function createNodeAction(designer) {
  return (type, from) => {
    let node;
    switch (type) {
      case "Text":
        node = TextNode.create(designer);
        break;
      case "Image":
        node = ImageNode.create(designer);
        break;
      case "Line":
        node = LineNode.create(designer);
        break;
      case "Barcode":
        node = BarcodeNode.create(designer, window.icons.barcode);
        break;
      case "Rect":
        node = RectNode.create(designer);
        break;
      case "Circle":
        node = CircleNode.create(designer);
        break;
      case "QrCode":
        node = QrCodeNode.create(designer, window.icons.qrcode);
        break;
    }

    if (from?.attrs) {
      node.updateAttrs(from.attrs);
    }

    designer.emit("select:node", node, node.getNodeAttrs());
  };
}
