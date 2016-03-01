if (location.pathname === "/") {
  boost(0);
}

document.addEventListener('DOMContentLoaded', function() {
  replaceAll('.date a', function(link) {
    return link.querySelector('time');
  });

  replaceAll('a.keyword', function(link) {
    return link.childNodes[0];
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
      var ga = node.id === "google-analytics-script";
      var self = node === selfScriptTag;

      return !self && !pushdog && !ga;
    });
    boostCount += 1;
    console.log(boostCount);
    if (boostCount < 20) {
      setTimeout(_boost, boostCount * 100);
    }
  }
}

function replaceAll(selector, fn) {
  toArray(document.querySelectorAll(selector)).forEach(function(node) {
    node.parentNode.replaceChild(fn(node), node);
  });
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
