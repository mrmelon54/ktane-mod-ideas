<script>
  import { onMount } from "svelte";
  import IdeaRow from "~/component/IdeaRow.svelte";
  import SearchBar from "~/component/SearchBar.svelte";

  export let fetch;
  let rows;

  onMount(() => {
    fetch()
      .then((x) => {
        console.log(x);
        rows = x;
      })
      .catch(() => {
        errorMessage = "Failed to load ideas";
      });
  });
</script>

<SearchBar />
<table class="full-width state-table">
  <td class="state-bg-green">Currently being implemented [🗸]</td>
  <td class="state-bg-yellow">Ready to be implemented [❖]</td>
  <td class="state-bg-red">In progress, not ready [✘]</td>
</table>
<table class="full-width">
  {#if rows === undefined}
    <IdeaRow idea={null} />
  {:else}
    {#each rows as row (row.id)}
      <IdeaRow idea={row} />
    {/each}
  {/if}
</table>

<style lang="scss">
  table.full-width {
    width: 100%;
  }

  table.state-table {
    border-collapse: collapse;
    &,
    tr,
    td {
      border: none;
    }
    td {
      padding: 4px;
    }
  }
</style>
