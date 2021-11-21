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

export const drawMeme = ({canvas, ctx, path, creditPrice='', creditDate='', meme_caption=''}) => {
  let imageObj = new Image();
  imageObj.src = path; 
  imageObj.onload = function(){

    // set bg
    ctx.fillStyle = '#fff';
    ctx.fillRect(0, 0, canvas.width, canvas.height);
    
    // load the image
    let autoHeight = canvas.width * imageObj.height / imageObj.width
    let y = canvas.height - autoHeight
    ctx.drawImage(imageObj, 0, y, canvas.width, autoHeight);

    // meme caption
    let fontSize
    if (meme_caption.length > 300) { fontSize = 12 } else
    if (meme_caption.length > 200) { fontSize = 18 } else
    if (meme_caption.length > 180) { fontSize = 22 } else
    if (meme_caption.length > 150) { fontSize = 24 } else
    if (meme_caption.length > 100) { fontSize = 32 } else
    if (meme_caption.length > 40) { fontSize = 40 } else
    if (meme_caption.length > 30) { fontSize = 54 } else
    if (meme_caption.length > 10) { fontSize = 60 } else
    { fontSize = 68 }
    let textMaxWidth = canvas.width-40;
    let lineHeight = fontSize + (fontSize*0.4);
    let textY = fontSize + 20;
    let textX = (canvas.width - textMaxWidth ) / 2;
    ctx.font = "Bold "+fontSize+"px 'Helvetica'";
    ctx.fillStyle = 'black';
    textArea(ctx, meme_caption, textX, textY, textMaxWidth, lineHeight);
    let crText = 'C'+'r'+'y'+'p'+'t'+'o'+'M'+'e'+'m'+'e'+'.'+'W'+'T'+'F';

    // credit text
    let strokedText = ({ctx, text, textX, textY}) => {
      ctx.font = "10px 'Arial'";
      ctx.strokeStyle = 'rgba(0,0,0,.3)';
      ctx.lineWidth = 3;
      ctx.strokeText(text, textX, textY);
      ctx.fillStyle = 'rgba(255,255,255,0.6)';
      ctx.fillText(text, textX, textY);
    }
    strokedText({
      ctx: ctx,
      text: creditDate,
      textX: 10,
      textY: canvas.height - 20,
    })
    strokedText({
      ctx: ctx,
      text: creditPrice,
      textX: 9,
      textY: canvas.height - 10,
    })
    strokedText({
      ctx: ctx,
      text: crText,
      textX: canvas.width - 90,
      textY: canvas.height - 10,
    })


  };
}