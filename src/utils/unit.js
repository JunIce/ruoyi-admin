export const unitConvert = {
  pxToMm: (px) => {
    // Standard conversion factor from pixels to millimeters at 96 DPI
    const PX_TO_MM = 25.4 / 96; // 1 inch = 25.4 mm, 1 inch = 96 px

    // Get the device pixel ratio (DPR)
    const dpr = window.devicePixelRatio || 1;

    // Adjust the pixel value based on the device pixel ratio
    const adjustedPx = px * dpr;

    // Convert the adjusted pixel value to millimeters
    return adjustedPx * PX_TO_MM;
  },
  mmToPx: (mm) => {
    // Standard conversion factor from millimeters to pixels at 96 DPI
    const MM_TO_PX = 96 / 25.4; // 1 inch = 25.4 mm, 1 inch = 96 px

    // Get the device pixel ratio (DPR)
    const dpr = window.devicePixelRatio || 1;

    // Convert the millimeter value to pixels
    let px = mm * MM_TO_PX;

    // Adjust the pixel value based on the device pixel ratio
    return px / dpr;
  },
};
