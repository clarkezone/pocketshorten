import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  vus: 100,
  duration: "30s",
};

function testSuccess() {
  let res = http.get("https://pocketshorten-stage.dev.clarkezone.dev/bp");
  check(res, {
    "status was 200": (r) => r.status === 200,
    "redirected correctly": (r) =>
      r.url === "https://blog-preview.clarkezone.dev",
  });
}

function testShortlinkNotFound() {
  let res = http.get("https://pocketshorten-stage.dev.clarkezone.dev/nf");
  check(res, {
    "status was 404": (r) => r.status === 404,
    "redirected correctly": (r) =>
      r.url === "https://pocketshorten-stage.dev.clarkezone.dev/nf",
  });
}

function testBadRequest() {
  let res = http.get("https://pocketshorten-stage.dev.clarkezone.dev");
  check(res, {
    "status was 404": (r) => r.status === 404,
    "redirected correctly": (r) =>
      r.url === "https://pocketshorten-stage.dev.clarkezone.dev",
  });
}

export default function () {
  testSuccess();
  sleep(1);
  testShortlinkNotFound();
  sleep(1);
  testBadRequest();
  sleep(1);
}
