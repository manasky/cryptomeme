const textArea = (context, text, x, y, maxWidth, lineHeight) => {
  let words = text.split(' ');
  let line = '';

  for(let n = 0; n < words.length; n++) {
    let testLine = line + words[n] + ' ';
    let metrics = context.measureText(testLine);
    let testWidth = metrics.width;
    if (testWidth > maxWidth && n > 0) {
      context.fillText(line, x, y);
      line = words[n] + ' ';
      y += lineHeight;
    }
    else {
      line = testLine;
    }
  }
  context.fillText(line, x, y);
}

const strokedText = ({ctx, text, textX, textY}) => {
  ctx.font = "10px 'Arial'";
  ctx.strokeStyle = 'rgba(0,0,0,.3)';
  ctx.lineWidth = 3;
  ctx.strokeText(text, textX, textY);
  ctx.fillStyle = 'rgba(255,255,255,0.6)';
  ctx.fillText(text, textX, textY);
}

export const drawMeme = ({
  canvas,
  ctx,
  bgImage,
  caption='',
  logo,
  footer_left,
  footer_right,
  meme_content,
}) => {

  // set bg
  ctx.fillStyle = '#fff';
  ctx.fillRect(0, 0, canvas.width, canvas.height);

  let footerHeight = 80;
  let logoSize = 50;
  let footerPaddingX = 20;
  let footerPaddingY = canvas.height-footerHeight+(footerHeight/2)-(logoSize/2);

  let imageObj = new Image();
  imageObj.src = bgImage; 
  imageObj.onload = function(){

    // set bg
    ctx.fillStyle = '#fff';
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    // load the image
    var hRatio = canvas.width / imageObj.width    ;
    var vRatio = (canvas.height-footerHeight) / imageObj.height  ;
    var ratio  = Math.min ( hRatio, vRatio );

    let imgW = imageObj.width * ratio
    let imgH = imageObj.height * ratio
    let imgX = (canvas.width - imgW)/2
    let imgY = canvas.height -footerHeight - imgH
    
    ctx.drawImage(imageObj, imgX, imgY, imgW, imgH);

    // meme caption
    let fontSize
    if (caption.length > 400) { fontSize = 18 } else
    if (caption.length > 250) { fontSize = 23 } else
    if (caption.length > 150) { fontSize = 24 } else
    if (caption.length > 100) { fontSize = 32 } else
    if (caption.length > 40) { fontSize = 40 } else
    if (caption.length > 30) { fontSize = 54 } else
    if (caption.length > 10) { fontSize = 70 } else
    { fontSize = 90 }
    let textMaxWidth = canvas.width-40;
    let lineHeight = fontSize + (fontSize*0.1);
    let textY = fontSize + 20;
    let textX = (canvas.width - textMaxWidth ) / 2;
    ctx.font = "Bold "+fontSize+"px 'Helvetica'";
    ctx.fillStyle = 'black';
    textArea(ctx, caption, textX, textY, textMaxWidth, lineHeight);
    let crText = 'C'+'r'+'y'+'p'+'t'+'o'+'M'+'e'+'m'+'e'+'.'+'W'+'T'+'F';

    // meme content
    if (meme_content && meme_content.length > 0) {
      meme_content.forEach(item => {
        if (item.image) {
          let contentImgObj = new Image();
          contentImgObj.src = item.image; 
          contentImgObj.onload = function(){
            ctx.drawImage(contentImgObj, item.x, item.y, 100, 100);
          }
        }else if(item.text){
          ctx.font = "Bold 40px 'Arial'";
          ctx.strokeStyle = 'white';
          ctx.lineWidth = 8;
          ctx.strokeText(item.text, item.x, item.y);
          ctx.fillStyle = 'black';
          ctx.fillText(item.text, item.x, item.y);
        }
      })
    }

    strokedText({
      ctx: ctx,
      text: crText,
      textX: canvas.width - 90,
      textY: canvas.height - footerHeight - 10,
    })
    
    // set footer bg
    // ctx.fillStyle = '#2a2e37';
    ctx.fillStyle = '#000';
    ctx.fillRect(0, canvas.height-footerHeight, canvas.width, canvas.height);
    let logoObj = new Image();
    // logoObj.setAttribute('crossorigin', 'anonymous');
    logoObj.src = logo; 
    logoObj.onload = function(){
      ctx.drawImage(logoObj, footerPaddingX, footerPaddingY, logoSize, logoSize);
    }
    // footer left
    ctx.font = "Bold 20px 'Arial'";
    ctx.fillStyle = '#fff';
    ctx.fillText(footer_left.title, footerPaddingX + logoSize + 10, footerPaddingY + 22);

    ctx.font = "14px 'Andale Mono', 'ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', 'Liberation Mono', 'Courier New', 'monospace'";
    ctx.fillStyle = '#ccc';
    ctx.fillText(footer_left.desc, footerPaddingX + logoSize + 10, footerPaddingY + 38);

    // footer right
    ctx.font = "20px 'Andale Mono', 'ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', 'Liberation Mono', 'Courier New', 'monospace'";
    ctx.fillStyle = '#fff';
    ctx.fillText(footer_right.title, canvas.width - ctx.measureText(footer_right.title).width - footerPaddingX, footerPaddingY + 22);

    ctx.font = "14px 'Andale Mono', 'ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', 'Liberation Mono', 'Courier New', 'monospace'";
    ctx.fillStyle = footer_right.color;
    ctx.fillText(footer_right.desc, canvas.width - ctx.measureText(footer_right.desc).width - footerPaddingX, footerPaddingY + 38);
  };

}