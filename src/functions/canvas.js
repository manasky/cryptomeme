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

export const bgBox = ({canvas, ctx, color = '#fafafa'}) => {
  ctx.fillStyle = color;
  ctx.fillRect(0, 0, canvas.width, canvas.height);
}

export const bottomImage = ({canvas, ctx, path}) => {
  let imageObj = new Image();
  imageObj.src = path; 
  imageObj.onload = function(){
    let autoHeight = canvas.width * imageObj.height / imageObj.width
    let y = canvas.height - autoHeight
    ctx.drawImage(imageObj, 0, y, canvas.width, autoHeight);
  };
}

export const topText = ({canvas, ctx, color= 'black', text, fontSize}) => {
  if (text.length > 300) { fontSize = 12 } else
  if (text.length > 200) { fontSize = 18 } else
  if (text.length > 180) { fontSize = 22 } else
  if (text.length > 150) { fontSize = 24 } else
  if (text.length > 100) { fontSize = 28 } else
  if (text.length > 40) { fontSize = 36 } else
  if (text.length > 30) { fontSize = 42 } else
  if (text.length > 10) { fontSize = 60 } else
  { fontSize = 68 }

  let textMaxWidth = canvas.width-40;
  let lineHeight = fontSize + (fontSize*0.4);
  let textY = fontSize + 20;
  let textX = (canvas.width - textMaxWidth ) / 2;
  ctx.font = "Bold "+fontSize+"px 'Helvetica'";
  ctx.fillStyle = color;
  textArea(ctx, text, textX, textY, textMaxWidth, lineHeight);
}