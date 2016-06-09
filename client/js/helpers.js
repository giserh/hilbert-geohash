
export function getUrlParams(splitCharNo){
  return [window.location.pathname.substring(0,splitCharNo), window.location.pathname.substring(splitCharNo)];
}

export function selectElement(elementId){
  return document.getElementById(elementId);
}

export function pushState(url, title){
  return window.history.pushState({},title,url);
}

export function setPushStateListener(fn) {
  return window.addEventListener('popstate', fn);
};

export function getHostName(){
  return window.location.href.split('/').splice(0,3).join('/');
}
