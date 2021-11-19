export const download = (canvasEl, name = 'filename' ) => {
  console.log('test')
  document.getElementById(canvasEl).download = "image.png";
  window.location.href=document.getElementById(canvasEl).toDataURL("image/png").replace("image/png", "image/octet-stream");
}