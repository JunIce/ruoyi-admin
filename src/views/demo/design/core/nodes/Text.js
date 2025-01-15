import Konva from 'konva';
import { AbstractNode } from './AbstractNode';

export class TextNode extends AbstractNode {
  type = 'Text';

  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      width: 30,
      height: 20,
      fontSize: 8,
    };
  }
  
  static create(designer) {
    let text = new Konva.Text({
      ...this.commonAttrs,
      wrap: 'word',
      draggable: true,
      resize: true,
      draggable: true,
      className: "Text",
    });

    let node = new TextNode(text, designer);

    node.draw(designer.layer);

    return node;
  }

  get columns() {
    return [
      {
        label: 'Width',
        prop: 'width',
        type: 'inputNumber',
      },
      {
        label: 'Height',
        prop: 'height',
        type: 'inputNumber',
      },
      {
        label: 'X',
        prop: 'x',
        type: 'inputNumber',
      },
      {
        label: 'Y',
        prop: 'y',
        type: 'inputNumber',
      },
      {
        label: 'ScaleX',
        prop: 'scaleX',
        type: 'inputNumber',
      },
      {
        label: 'ScaleY',
        prop: 'scaleY',
        type: 'inputNumber',
      },
      {
        label: '文本',
        prop: 'text',
        type: 'input',
      },
      {
        label: '字体大小',
        prop: 'fontSize',
        type: 'inputNumber',
        trigger: 'input',
      },
      {
        label: '绑定属性',
        prop: 'bindProp',
        type: 'input',
      },
    ];
  }
}
