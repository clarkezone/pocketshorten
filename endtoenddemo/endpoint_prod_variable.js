import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "10s", target: 10 },
    { duration: "5s", target: 100 },
    { duration: "20s", target: 200 },
    { duration: "100s", target: 200 },
    { duration: "500s", target: 10 },
  ],
};

function testSuccess() {
  const sourceUrls = [
    "https://psdemo.clarkezone.dev/tsh",
    "https://psdemo.clarkezone.dev/lol",
    "https://psdemo.clarkezone.dev/m2",
    "https://psdemo.clarkezone.dev/m3",
  ];
  const targetUrls = [
    "https://psdemotarget.clarkezone.dev/",
    "https://psdemotarget.clarkezone.dev/meme1.html",
    "https://psdemotarget.clarkezone.dev/meme2.html",
    "https://psdemotarget.clarkezone.dev/meme3.html",
  ];
  const randomIndex = Math.floor(Math.random() * sourceUrls.length);
  const url = sourceUrls[randomIndex];

  let res = http.get(url);
  // console.log(res.url + targetUrls[randomIndex] + res.status);
  check(res, {
    "status was 200": (r) => r.status === 200,
    "redirected correctly": (r) => r.url === targetUrls[randomIndex],
  });
}

export default function () {
  testSuccess();
}
