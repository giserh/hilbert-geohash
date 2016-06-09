import { getHash, getPoint } from './api';
import { loadMap, setMarker, setZoom, addRectangle, removeHashPath } from './map';
import { pushState, getHostName } from './helpers';

let hashField, hashInput,pointInput, pointField, pathButton, hashValue;

let setInfo = (res) =>{
  let { hash, point } = res;
  hashField.innerHTML = getHostName()+'/h/'+hash;
  hashInput.value = hash;
  pointInput.value = `${point.lat},${point.lng}`;
  hashField.href = '/h/'+hash;
  pointField.innerHTML = getHostName()+`/p/${point.lat},${point.lng}`;
  pointField.href = `/p/${point.lat},${point.lng}`;
  hashValue = hash;
};

let updateInfo = (res) =>{
  let { hash, point } = res;
  hashField.innerHTML = getHostName()+'/h/'+hash;
  hashField.href = '/h/'+hash;
  pointField.innerHTML = getHostName()+`/p/${point.lat},${point.lng}`;
  pointField.href = `/p/${point.lat},${point.lng}`;
  showHashPath(hash);
  pushState('/h/'+hash, 'Hash:'+hash);
  hashValue = hash;
};

export function init(hi,pi,hf,pf,pb,map){
  hashField = hf;
  pointField = pf;
  hashInput = hi;
  pointInput = pi;
  pathButton = pb;
  return loadMap(map).then(gmap=>{
    gmap.addListener('click', e => {
      getHash(e.latLng.lat(),e.latLng.lng())
      .then(res=>{
        setMarker(res.point);
        setInfo(res);
      });
    });
    hashInput.addEventListener('input', e => {
      getPoint(e.srcElement.value)
      .then(res=>{
        setMarker(res.point);
        updateInfo(res);
      });
    });
    pointInput.addEventListener('keyup', e => {
      if (e.keyCode != 13) return;
      let [lat,lng] = e.srcElement.value.split(',');
      getHash(lat,lng)
      .then(res=>{
        setMarker(res.point);
        updateInfo(res);
      });
    });

    pathButton.addEventListener('click', e=>{
      showHashPath(hashValue)
    });
  });
};

export function showHash(hash) {
  removeHashPath();
  getPoint(hash)
    .then(res=>{
      setMarker(res.point);
      setInfo(res);
      setZoom(res.point);
      pushState(`/h/${res.hash}`,`Hash:${res.hash}`);
    });
}

export function showHashPath(hash){
  removeHashPath();
  if(!pathButton.checked) return;
  for (let i=0;i<hash.length;i++){
    getPoint(hash.substring(0,i+1)).then(res=>addRectangle(res.point));
  }
}

export function showPoint(lat,lng) {
  getHash(lat,lng)
    .then(res=>{
      setMarker(res.point);
      setInfo(res);
      setZoom(res.point);
      pushState(`/p/${lat},${lng}`,`Point:${lat},${lng}`);
    });
}
