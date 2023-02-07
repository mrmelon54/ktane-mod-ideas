<script>
  let query = "";

  function searchButton() {}

  function clearButton() {}
</script>

<div class="search-bar">
  <label for="query"><input bind:value={query} type="text" /></label>
  <button on:click={searchButton}>Search</button>
  <button on:click={clearButton}>Clear</button>
</div>
