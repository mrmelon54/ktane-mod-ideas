<script lang="ts">
  import { API_URL } from "~/internal/api";
  import { clickOutside } from "~/lib/clickOutside";
  import { popupCenterScreen } from "~/utils/window";

  export let closeModal: () => void;

  function loginWithDiscordStage2() {
    closeModal();
    popupCenterScreen(`${API_URL}/login?in_popup`, "Login", 600, 900, false);
  }
</script>

<div class="login-modal">
  <div
    class="modal-content"
    use:clickOutside
    on:click_outside={() => closeModal()}
  >
    <header>
      <h2>
        Login with Discord
        <button on:click={() => closeModal()} />
      </h2>
    </header>

    <div class="modal-container">
      <p>
        By logging into this website you give permission for your Discord ID and
        Discord tag to be stored in a database so your tag can be displayed
        along side any ideas you add. Your Discord ID is also linked to each
        vote you make to ensure you don't vote multiple times, however this is
        not displayed on the website so only you know what ideas you have voted
        on.
      </p>
      <p>
        Administrators of this site are able to ban your Discord account from
        using the site at anytime if you are not using the site for its intended
        purpose, you are breaking the voting system or you are putting adult
        content on the site. This list is not in anyway a complete list of
        bannable offenses and the administrators are allowed to ban your account
        for any reason.
      </p>
      <p>Enough with the legal nonsense... click below to login.</p>
    </div>

    <footer class="modal-container">
      <button class="modal-button" on:click={() => loginWithDiscordStage2()}>
        Login with Discord
      </button>
    </footer>
  </div>
</div>

<style lang="scss">
  .login-modal {
    font-family: "Roboto", sans-serif;
    background-color: rgba(0, 0, 0, 0.5);
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;

    .modal-content {
      position: absolute;
      top: 60px;
      left: 50px;
      right: 50px;
      background: #242424;
      border-radius: 12px;
      overflow: hidden;

      header {
        background: #5865f2;

        h2 {
          margin: 0;
          padding: 16px 32px;
          line-height: 40px;

          button {
            float: right;
            aspect-ratio: 1/1;
            height: 40px;
            border: none;
            background: #dd5858;
            cursor: pointer;
            user-select: none;
            position: relative;
            border-radius: 8px;

            &::before {
              content: "×";
              font-size: 40px;
              position: absolute;
              top: 50%;
              left: 50%;
              transform: translate(-50%, -50% - 5);
            }
          }
        }
      }

      .modal-container {
        padding: 16px;
      }

      .modal-button {
        background: #5865f2;
        color: #ffffff;
        border: 0;
        font-size: 1em;
        padding: 12px;
        border-radius: 8px;
        cursor: pointer;
      }
    }
  }
</style>
