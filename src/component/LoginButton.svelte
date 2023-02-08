<script>
  import { onMount } from "svelte";
  import { popupCenterScreen } from "~/utils/window";
  import { user } from "~/internal/stores";
  import { API_URL } from "~/internal/api";
  import { isObject } from "~/utils/utils";

  onMount(() => {
    check_user();
  });

  function handleLoginClick() {
    popupCenterScreen(`${API_URL}/login?in_popup`, "Login", 600, 900, false);
  }

  function handleLogoutClick() {
    popupCenterScreen(`${API_URL}/logout?in_popup`, "Logout", 600, 900, false);
    user.set(null);
  }

  function check_user() {
    const f = document.createElement("iframe");
    f.src = `${API_URL}/check`;
    f.style.display = "none";
    document.body.appendChild(f);
  }

  window.onmessage = function (event) {
    if (event.origin !== API_URL) return;
    if (isObject(event.data)) {
      console.log(event.data);
      if (isObject(event.data.user)) {
        let d = Object.assign(
          {
            id: null,
            discord_name: null,
            discord_discriminator: null,
            picture: null,
            banned: false,
            admin: false,
          },
          event.data.user
        );
        if (d.id === null) {
          alert(
            "Failed to log user in: the login data is structured correctly but probably corrupted"
          );
          return;
        }
        user.set(d);
        return;
      }
    }
    alert("Failed to log user in: the login data was probably corrupted");
  };
</script>

{#if $user}
  <div id="loginMenu">
    <button id="loginProfile" class="account-btn rounded-blue">
      <img id="loginMenuAvatar" src={$user.picture} alt="" />
      <span id="loginMenuName">{$user.discord_name}</span>
    </button>
    <div id="loginDropdown">
      <a href="/admin/dashboard" class="dropdown-pill">Dashboard</a>
      <a href="/admin/announcements" class="dropdown-pill">Announcements</a>
      <a href="/admin/settings" class="dropdown-pill">Settings</a>
      <button class="dropdown-pill" on:click={handleLogoutClick}>Logout</button>
    </div>
  </div>
{:else}
  <button
    id="loginBtn"
    class="account-btn rounded-blue"
    on:click={handleLoginClick}>Login</button
  >
{/if}

<style>
  #loginMenu {
    position: relative;
  }

  #loginMenu #loginProfile {
    position: relative;
    padding-left: 44px;
  }

  button.account-btn#loginProfile img#loginMenuAvatar {
    position: absolute;
    left: 2px;
    width: 32px;
    top: 2px;
    border-radius: 8px;
  }

  #loginMenu #loginDropdown {
    display: none;
  }

  #loginMenu:hover #loginDropdown {
    position: absolute;
    top: 100%;
    right: 0;
    display: flex;
    flex-direction: column;
    flex: 1 1 100%;
    background: #35425c;
    border-radius: 8px;
    z-index: 999;
    padding: 10px;
  }

  #loginMenu #loginDropdown .dropdown-pill {
    margin: 2px;
    padding: 6px;
    display: inline-block;
    cursor: pointer;
  }

  #loginMenu #loginDropdown .dropdown-pill:hover {
    background: #55555555;
    border-radius: 8px;
  }
</style>
