import{S as C,i as U,s as F,e as A,c as S,a as H,d as a,b as m,f as h,R as I,D as j,A as O,T as R,Q as X,t as Y,k as D,g as z,n as E,G as q,h as B}from"./vendor-50905681.js";let v=t=>t>1?t.toLocaleString("en-US",{maximumFractionDigits:2}):t.toLocaleString("en-US",{maximumFractionDigits:10}),x=t=>{let i=Math.min(...t);for(var l=0;l<t.length;l++)t[l]=t[l]-i;return t},p=(t,i)=>Object.keys(t).map(l=>t[l][i]);const G=(t,i,l,c,u,r)=>{let n=i.split(" "),T="";for(let w=0;w<n.length;w++){let o=T+n[w]+" ";t.measureText(o).width>u&&w>0?(t.fillText(T,l,c),T=n[w]+" ",c+=r):T=o}t.fillText(T,l,c)},N=({canvas:t,ctx:i,path:l,creditPrice:c="",creditDate:u="",meme_caption:r=""})=>{let n=new Image;n.src=l,n.onload=function(){i.fillStyle="#fff",i.fillRect(0,0,t.width,t.height);let T=t.width*n.height/n.width,w=t.height-T;i.drawImage(n,0,w,t.width,T);let o;r.length>300?o=12:r.length>200?o=18:r.length>180?o=22:r.length>150?o=24:r.length>100?o=32:r.length>40?o=40:r.length>30?o=54:r.length>10?o=60:o=68;let d=t.width-40,M=o+o*.4,b=o+20,g=(t.width-d)/2;i.font="Bold "+o+"px 'Helvetica'",i.fillStyle="black",G(i,r,g,b,d,M);let s="CryptoMeme.WTF",y=({ctx:_,text:L,textX:k,textY:P})=>{_.font="10px 'Arial'",_.strokeStyle="rgba(0,0,0,.3)",_.lineWidth=3,_.strokeText(L,k,P),_.fillStyle="rgba(255,255,255,0.6)",_.fillText(L,k,P)};y({ctx:i,text:u,textX:10,textY:t.height-20}),y({ctx:i,text:c,textX:9,textY:t.height-10}),y({ctx:i,text:s,textX:t.width-90,textY:t.height-10})}};function Q(t){let i,l,c,u;return{c(){i=A("canvas"),this.h()},l(r){i=S(r,"CANVAS",{class:!0,width:!0,height:!0,id:!0}),H(i).forEach(a),this.h()},h(){m(i,"class",l=""+(t[2]+" -my-3 -mr-4 border border-base-content border-opacity-5 cursor-pointer")),m(i,"width","700"),m(i,"height","600"),m(i,"id",t[0])},m(r,n){h(r,i,n),t[10](i),c||(u=I(i,"click",function(){R(t[1]?t[4]():"")&&(t[1]?t[4]():"").apply(this,arguments)}),c=!0)},p(r,[n]){t=r,n&4&&l!==(l=""+(t[2]+" -my-3 -mr-4 border border-base-content border-opacity-5 cursor-pointer"))&&m(i,"class",l),n&1&&m(i,"id",t[0])},i:j,o:j,d(r){r&&a(i),t[10](null),c=!1,u()}}}function V(t,i,l){let{id:c}=i,{meme:u}=i,{coin_name:r}=i,{meme_caption:n}=i,{creditPrice:T=""}=i,{creditDate:w=""}=i,{clickToDownload:o=!1}=i,{classes:d="w-52 rounded"}=i,M,b;O(()=>{let s=M.getContext("2d");N({canvas:M,ctx:s,path:u,meme_caption:n,creditPrice:T,creditDate:w}),l(4,b=()=>{let y=document.createElement("a");y.download=r+"-meme-"+new Date(Date.now()).toLocaleString("en-US",{year:"numeric",month:"short",day:"2-digit"})+" [cryptomeme.wtf].jpg",y.href=M.toDataURL(),y.click()})});function g(s){X[s?"unshift":"push"](()=>{M=s,l(3,M)})}return t.$$set=s=>{"id"in s&&l(0,c=s.id),"meme"in s&&l(5,u=s.meme),"coin_name"in s&&l(6,r=s.coin_name),"meme_caption"in s&&l(7,n=s.meme_caption),"creditPrice"in s&&l(8,T=s.creditPrice),"creditDate"in s&&l(9,w=s.creditDate),"clickToDownload"in s&&l(1,o=s.clickToDownload),"classes"in s&&l(2,d=s.classes)},[c,o,d,M,b,u,r,n,T,w,g]}class $ extends C{constructor(i){super();U(this,i,V,Q,F,{id:0,meme:5,coin_name:6,meme_caption:7,creditPrice:8,creditDate:9,clickToDownload:1,classes:2})}}function J(t){let i,l,c,u,r,n,T,w,o,d,M,b,g,s,y,_,L,k,P,W;return{c(){i=A("title"),l=Y(t[0]),c=D(),u=A("meta"),r=D(),n=A("meta"),T=D(),w=A("meta"),o=D(),d=A("meta"),M=D(),b=A("meta"),g=D(),s=A("meta"),y=D(),_=A("meta"),L=D(),k=A("meta"),P=D(),W=A("meta"),this.h()},l(e){i=S(e,"TITLE",{});var f=H(i);l=z(f,t[0]),f.forEach(a),c=E(e),u=S(e,"META",{name:!0,content:!0}),r=E(e),n=S(e,"META",{name:!0,content:!0}),T=E(e),w=S(e,"META",{name:!0,content:!0}),o=E(e),d=S(e,"META",{name:!0,content:!0}),M=E(e),b=S(e,"META",{name:!0,content:!0}),g=E(e),s=S(e,"META",{name:!0,content:!0}),y=E(e),_=S(e,"META",{property:!0,content:!0}),L=E(e),k=S(e,"META",{property:!0,content:!0}),P=E(e),W=S(e,"META",{property:!0,content:!0}),this.h()},h(){m(u,"name","description"),m(u,"content",t[1]),m(n,"name","twitter:card"),m(n,"content","summary"),m(w,"name","twitter:title"),m(w,"content",t[0]),m(d,"name","twitter:description"),m(d,"content",t[1]),m(b,"name","twitter:image"),m(b,"content",t[2]),m(s,"name","twitter:image:alt"),m(s,"content",t[0]),m(_,"property","og:title"),m(_,"content",t[0]),m(k,"property","og:description"),m(k,"content",t[1]),m(W,"property","og:image"),m(W,"content",t[2])},m(e,f){h(e,i,f),q(i,l),h(e,c,f),h(e,u,f),h(e,r,f),h(e,n,f),h(e,T,f),h(e,w,f),h(e,o,f),h(e,d,f),h(e,M,f),h(e,b,f),h(e,g,f),h(e,s,f),h(e,y,f),h(e,_,f),h(e,L,f),h(e,k,f),h(e,P,f),h(e,W,f)},p(e,[f]){f&1&&B(l,e[0]),f&2&&m(u,"content",e[1]),f&1&&m(w,"content",e[0]),f&2&&m(d,"content",e[1]),f&4&&m(b,"content",e[2]),f&1&&m(s,"content",e[0]),f&1&&m(_,"content",e[0]),f&2&&m(k,"content",e[1]),f&4&&m(W,"content",e[2])},i:j,o:j,d(e){e&&a(i),e&&a(c),e&&a(u),e&&a(r),e&&a(n),e&&a(T),e&&a(w),e&&a(o),e&&a(d),e&&a(M),e&&a(b),e&&a(g),e&&a(s),e&&a(y),e&&a(_),e&&a(L),e&&a(k),e&&a(P),e&&a(W)}}}function K(t,i,l){let{title:c}=i,{desc:u}=i,{img:r}=i;return t.$$set=n=>{"title"in n&&l(0,c=n.title),"desc"in n&&l(1,u=n.desc),"img"in n&&l(2,r=n.img)},[c,u,r]}class ee extends C{constructor(i){super();U(this,i,K,J,F,{title:0,desc:1,img:2})}}export{$ as M,ee as S,p as e,v as f,x as n};
