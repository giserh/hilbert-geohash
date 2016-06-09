import request from 'superagent';
import alertify from 'alertify.js';

export function getPoint(hash) {
  return new Promise(resolve=>{
    request.post('/api/hash')
      .send({hash})
      .end((err,res)=>err ? alertify.error("Invalid hash"):resolve(res.body));
  });
}


export function getHash(lat,lng) {
  lat = typeof lat == "string" ? Number(lat):lat;
  lng = typeof lng == "string" ? Number(lng):lng;
  return new Promise(resolve=>{
    request.post('/api/point')
      .send({lat,lng})
      .end((err,res)=>err ? alertify.error("Invalid point"):resolve(res.body));
  });
}
