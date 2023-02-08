<script lang="ts">
  let query = "";

  function searchButton() {}

  function clearButton() {
    query = "";
  }
</script>

<div class="search-bar">
  <label for="query"><input bind:value={query} type="text" /></label>
  <button on:click={searchButton}>Search</button>
  <button on:click={clearButton}>Clear</button>
</div>

<style lang="scss">
  .search-bar {
    padding: 8px;
    display: flex;
    align-items: stretch;
    gap: 8px;

    @media screen and (prefers-color-scheme: light) {
      background-color: #bdbdbd;
    }

    input {
      padding: 4px;
      font-size: 1.1em;
    }

    button {
    }
  }
</style>
