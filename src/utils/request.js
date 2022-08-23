import instance from "./http";

const axios = ({ method, url, data, config }) => {
  method = method.toLowerCase();
  if (method === "post") {
    return instance.post(url, data, { ...config });
  } else if (method === "get") {
    return instance.get(url, { params: data, ...config });
  } else if (method === "delete") {
    return instance.delete(url, { params: data, ...config });
  } else if (method === "put") {
    return instance.put(url, data, { ...config });
  } else {
    console.error("UnKnown Method:" + method);
    return false;
  }
};

export default axios;
