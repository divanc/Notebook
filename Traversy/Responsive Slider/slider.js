const slides = document.querySelectorAll(".slide");
const next = document.querySelector("#next");
const prev = document.querySelector("#prev");

const auto = true;
const intervalTime = 2000;
let sliderInterval;

const nextSlide = () => {
  const current = document.querySelector(".current");
  current.classList.remove("current");

  if (current.nextElementSibling)
    current.nextElementSibling.classList.add("current");
  else slides[0].classList.add("current");

  setTimeout(() => current.classList.remove("current"));
};

const prevSlide = () => {
  const current = document.querySelector(".current");
  current.classList.remove("current");

  if (current.previousElementSibling) {
    current.previousElementSibling.classList.add("current");
  } else {
    slides[slides.length - 1].classList.add("current");
  }

  setTimeout(() => current.classList.remove("current"));
};

next.addEventListener("click", event => {
  nextSlide();
  if (auto) {
    clearInterval(sliderInterval);
    sliderInterval = setInterval(nextSlide, intervalTime);
  }
});
prev.addEventListener("click", event => {
  prevSlide();
  if (auto) {
    clearInterval(sliderInterval);
    sliderInterval = setInterval(nextSlide, intervalTime);
  }
});

if (auto) sliderInterval = setInterval(nextSlide, intervalTime);
