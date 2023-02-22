import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "100s", target: 10 },
    { duration: "5s", target: 100 },
    { duration: "20s", target: 200 },
    { duration: "100s", target: 200 },
    { duration: "500s", target: 10 },
  ],
};

function testSuccess() {
  let res = http.get("https://s.clarkezone.dev/bp");
  check(res, {
    "status was 200": (r) => r.status === 200,
    "redirected correctly": (r) =>
      r.url === "https://blog-preview.clarkezone.dev",
  });
}

function testShortlinkNotFound() {
  let res = http.get("https://s.clarkezone.dev?shortlink=nf");
  check(res, {
    "status was 404": (r) => r.status === 404,
    "redirected correctly": (r) =>
      r.url === "https://s.clarkezone.dev?shortlink=nf",
  });
}

export default function () {
  testSuccess();
  sleep(1);
  testShortlinkNotFound();
  sleep(1);
  testSuccess();
  sleep(1);
}
