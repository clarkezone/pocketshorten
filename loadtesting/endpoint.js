import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  vus: 10,
  duration: "30s",
};

export default function () {
  let res = http.get(
    "https://pocketshorten-stage.dev.clarkezone.dev?shortlink=tm"
  );
  check(res, {
    "status was 301": (r) => r.status === 301,
    "redirected correctly": (r) => r.url === "https://techmeme.com",
  });
  sleep(1);
}
