import GoogleMapsLoader from 'google-maps';

let GOOGLE;
const GAPIKEY = 'AIzaSyBkHAmahy4dhj9a4U6iAxqStKkxzd6mvL0';
let loadGoogle = GoogleApiKey => {
  GoogleMapsLoader.KEY = GoogleApiKey || GAPIKEY;
  return new Promise(res=>{
    GoogleMapsLoader.load(google=>{
      GOOGLE = google;
      return res();
    });
  });
};

let gmap;
export function loadMap(element) {
  return loadGoogle().then(()=>{
    gmap = new GOOGLE.maps.Map(element, {
     mapTypeId: GOOGLE.maps.MapTypeId.ROADMAP,
     draggableCursor:'crosshair',
    });
    return gmap;
  });
};

const ZOOM = 2;
export function setZoom(point) {
  let zoomBounds = new GOOGLE.maps.LatLngBounds(
    new google.maps.LatLng(point.lat - ZOOM*point.latErr,point.lng - ZOOM*0.5*point.lngErr),
    new google.maps.LatLng(point.lat + ZOOM*point.latErr, point.lng + ZOOM*0.5*point.lngErr)
  );
  return gmap.fitBounds(zoomBounds);
};

let rectangles = [];
export function addRectangle(point) {
  let rectangleBounds = new GOOGLE.maps.LatLngBounds(
    new google.maps.LatLng(point.lat - point.latErr,point.lng - point.lngErr),
    new google.maps.LatLng(point.lat + point.latErr, point.lng + point.lngErr)
  );
  rectangles.push(new GOOGLE.maps.Rectangle({
    strokeColor: '#FF0000',
    strokeOpacity: 0.8,
    strokeWeight: 2,
    fillColor: '#FF0000',
    fillOpacity: 0.0, //Trasnsparent fill
    map: gmap,
    clickable: false,
    bounds:rectangleBounds,
  }));
}

export function removeHashPath() {
  rectangles.forEach(rectangle=>{
    rectangle.setMap(null);
    rectangle = null;
  });
  rectangles = [];
}

let marker;
export function setMarker(position) {
  if (marker) {
    marker.setPosition(position)
    gmap.panTo(marker.getPosition());
    return marker;
  }
  marker = new GOOGLE.maps.Marker({
    position,
    map: gmap,
  });
  gmap.panTo(marker.getPosition());
  return marker
};
