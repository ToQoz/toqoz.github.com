if (location.pathname === "/") {
  boost(0);
}

document.addEventListener('DOMContentLoaded', function() {
  toArray(document.querySelectorAll('.date')).forEach(function(date) {
    var time = date.querySelector('time');
    var link = date.querySelector('a');
    date.removeChild(link);
    date.appendChild(time);
  });
});

function boost() {
  var selfScriptTag = document.currentScript;
  var boostCount = 0;
  _boost();

  function _boost() {
    removeAll('iframe');
    removeAll('script', function(node) {
      var src = node.getAttribute('src');
      var pushdog = src && src.indexOf('pushdog') !== -1;
      var self = node === selfScriptTag;
      return !self && !pushdog ;
    });
    boostCount += 1;
    console.log(boostCount);
    if (boostCount < 20) {
      setTimeout(_boost, boostCount * 100);
    }
  }
}

function removeAll(selector, test) {
  toArray(document.querySelectorAll(selector)).forEach(function(node) {
    if (!test || test(node)) {
      node.parentNode.removeChild(node);
    }
  });
}

function toArray(ary) {
  return Array.prototype.slice.apply(ary);
}
