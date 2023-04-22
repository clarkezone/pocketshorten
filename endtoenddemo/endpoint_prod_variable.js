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
  let res = http.get("https://psdemo.clarkezone.dev/tsh");
  check(res, {
    "status was 200": (r) => r.status === 200,
    "redirected correctly": (r) =>
      r.url === "https://psdemotarget.clarkezone.dev",
  });
}

export default function () {
  testSuccess();
}
