<script>
  import { onMount } from "svelte";
  import LoginButton from "~/component/LoginButton.svelte";
  import { fetchHomeResult } from "~/internal/api";
  import Loading from "~/component/Loading.svelte";
  import IdeaList from "~/component/IdeaList.svelte";
  import IdeaRow from "~/component/IdeaRow.svelte";

  let result;
  let errorMessage;
  onMount(async () => {
    fetchHomeResult()
      .then((v) => {
        result = v;
      })
      .catch(() => {
        errorMessage = "Failed to load ideas for homepage";
      });
  });
</script>

<div class="app-container">
  <header>
    <a href="/">
      <h1>KTaNE Mod Ideas</h1>
    </a>
    <LoginButton />
  </header>
  <main>
    {#if result === undefined}
      {#if errorMessage}
        <IdeaRow idea={null} />
      {:else}
        <div id="main-loading">
          <Loading />
        </div>
      {/if}
    {:else}
      <IdeaList fetch={fetchHomeResult} />
    {/if}
  </main>
</div>

<style lang="scss">
  .app-container {
    min-width: 1000px;
    max-width: calc(100vw - 32px);
    margin: auto;

    @media screen and (min-width: 1000px) {
      min-width: 1000px;
      max-width: calc(100vw - 32px);
      margin: auto;
    }

    header {
      padding: 0 60px;
      box-sizing: border-box;
      display: flex;
      align-items: center;
      justify-content: space-between;
      height: 80px;
      min-height: 80px;
      max-height: 80px;
      background-color: #2e323e;
      background-image: none;
      border-bottom-color: #8c827319;

      a {
        all: unset;

        h1 {
          font-size: 20px;
          font-weight: 500;
          line-height: 28px;
          color: #d0ccc6;
          cursor: pointer;
        }
      }
    }
  }
</style>
