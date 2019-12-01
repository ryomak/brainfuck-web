const getParam = parameterName => {
  var result = null,
    tmp = [];
  location.search
    .substr(1)
    .split("&")
    .forEach(function(item) {
      tmp = item.split("=");
      if (tmp[0] === parameterName) result = decodeURIComponent(tmp[1]);
    });
  return result;
};

const addParam = (key, value) => {
  return `${window.location.origin}/?${key}=${value}`;
};

export { getParam, addParam };
