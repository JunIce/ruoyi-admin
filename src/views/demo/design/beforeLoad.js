// 条码图片
import barcode from "./assets/barcode.png";
import qrcode from "./assets/qrcode.png";
import { loadImage } from "./core/util";

const icons = {
  barcode,
  qrcode,
};

function loadIcons() {
  window.icons = {};

  return Promise.all(
    Object.keys(icons).map((key) =>
      loadImage(icons[key]).then((image) => {
        window.icons[key] = image;
      })
    )
  );
}

export default { loadIcons };
