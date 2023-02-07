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
<table>
  <td><span class="state-bg-green">Currently being implemented [🗸]</span></td>
  <td><span class="state-bg-yellow">Ready to be implemented [❖]</span></td>
  <td><span class="state-bg-red">In progress, not ready [✘]</span></td>
</table>
<table>
  {#if rows === undefined}
    <IdeaRow idea={null} />
  {:else}
    {#each rows as row (row.id)}
      <IdeaRow idea={row} />
    {/each}
  {/if}
</table>
