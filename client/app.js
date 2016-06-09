"use strict";
import './styles/main.css';
import { selectElement as $, setPushStateListener, getUrlParams } from './js/helpers';
import { init, showHash, showPoint } from './js/gui';

let route = ()=> {
  let [path, param] = getUrlParams(3); //Gets the current browser URL splitted: Path and params
  switch (path) {
    case '/h/': //Url's starting with /h/ are hashes
      showHash(param); //Parmeter is a hash
      break;
    case '/p/': //Url's starting with /p/ are points
      let [lat,lng] = param.split(',');//Split params into lat,lng coordinates
      showPoint(lat,lng);
      break;
    default: //Everything else: Show home.
      showPoint(63.41678857590608,10.402773320674896);
  };
}
//Init the GUI.
init($('hashInput'),
  $('pointInput'),
  $('hash'),
  $('point'),
  $('pathButton'),
  $('map')).then(route);

// Update view when URL has changed.
setPushStateListener(route);
