export const formatValueToNumber = (obj, props) => {
  Object.keys(obj).forEach((k) => {
    if (props.includes(k)) {
      obj[k] = Number(obj[k]);
    }
  });

  return obj;
};

let map = {};
export const loadImage = (img) => {
  return new Promise((resolve, reject) => {
    if (map[img]) {
      return resolve(map[img]);
    }
    const image = new Image();
    image.onload = () => {
      map[img] = image;
      resolve(image);
    };
    image.onerror = reject;
    image.src = img;
    image.crossOrigin = 'Anonymous';
  });
};

export const pxUtils = {
  mmToPx,
  pxToMm,
};

export const DPI = (function estimateDPI() {
  // 创建一个隐藏的 div 元素用于测量
  const div = document.createElement("div");
  div.style.width = "1in";
  div.style.padding = "0";
  div.style.border = "none";
  div.style.margin = "0";
  document.body.appendChild(div);
  // 获取计算后的宽度（以像素为单位）
  const dpi = div.offsetWidth;
  document.body.removeChild(div);
  return dpi;
})();
console.log(`DPI: ${DPI}`);

// 毫米转像素
export function mmToPx(mm) {
  // Standard conversion factor from millimeters to pixels at 96 DPI
  const MM_TO_PX = DPI / 25.4; // 1 inch = 25.4 mm, 1 inch = 96 px

  // Get the device pixel ratio (DPR)
  const dpr = window.devicePixelRatio || 1;

  // Convert the millimeter value to pixels
  let px = mm * MM_TO_PX;

  // Adjust the pixel value based on the device pixel ratio
  return px / dpr;
}
// 像素转毫米
export function pxToMm(px) {
  // Standard conversion factor from pixels to millimeters at 96 DPI
  const PX_TO_MM = 25.4 / DPI; // 1 inch = 25.4 mm, 1 inch = 96 px

  // Get the device pixel ratio (DPR)
  const dpr = window.devicePixelRatio || 1;

  // Adjust the pixel value based on the device pixel ratio
  const adjustedPx = px * dpr;

  // Convert the adjusted pixel value to millimeters
  return adjustedPx * PX_TO_MM;
}
