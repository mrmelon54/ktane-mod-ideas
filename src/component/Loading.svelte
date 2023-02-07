<div class="loading-animation">
  <div class="pip" style="--angle: 0deg;" />
  <div class="pip" style="--angle: 60deg;" />
  <div class="pip" style="--angle: 120deg;" />
  <div class="pip" style="--angle: 180deg;" />
  <div class="pip" style="--angle: 240deg;" />
  <div class="pip" style="--angle: 300deg;" />
</div>

<style>
  .loading-animation {
    display: block;
    border: 8px solid #55aa55;
    border-radius: 50%;
    background: #aa5555;
    width: 100px;
    height: 100px;
    margin: auto;
    position: relative;
  }

  .loading-animation .pip {
    position: absolute;
    display: block;
    background: #555555;
    border-radius: 50%;
    top: calc(50% - 8px);
    left: calc(50% - 8px);
    width: 16px;
    height: 16px;
    --angle: 0deg;
    transform: rotate(var(--angle)) translate(32px);
    animation: loading-anim 1s infinite;
  }

  @keyframes loading-anim {
    0% {
      transform: rotate(var(--angle)) translate(32px);
    }

    100% {
      transform: rotate(calc(var(--angle) + 360deg)) translate(32px);
    }
  }
</style>
