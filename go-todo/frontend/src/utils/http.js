import axios from "axios";

const jwt =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTE4MTI2NTksImlkIjoyMH0.7cT5jXxulDnGze9QoqH4i8G_vXzeQit-kphsU2TtWx8";

const http = axios.create({
  baseURL: "http://localhost:8000/api/v1",
  headers: {
    Authorization: `Bearer ${jwt}`,
  },
});

// Return just the data object from the response
http.interceptors.response.use(({ data }) => {
  const { data: innerData, ...rest } = data;
  if (rest.Text === "OK") {
    return innerData && typeof innerData === "object"
      ? { ...rest, ...innerData } // eslint-ignore-line
      : data;
  }
  return data;
});

export { http };
