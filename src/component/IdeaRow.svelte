<script>
  export let idea;
  let firstLink = "https://google.com";
  $: firstLink = idea === null ? "" : idea.manual_url;
</script>

{#if idea === null}
  <tr>
    <td>No idea</td>
    <td />
    <td />
  </tr>
{:else}
  <tr>
    <td>
      <div>
        <a class="mod-name" target="_blank" href={firstLink}>{idea.name}</a>
        <i class="mod-author">by {idea.author}</i>
      </div>
      <div>{idea.description}</div>
      <div>{idea.notes}</div>
    </td>
    <td>{idea.vote}</td>
    {#if idea.state == 1}
      <td class="state-bg-red">✘</td>
    {:else if idea.state == 2}
      <td class="state-bg-yellow">❖</td>
    {:else if idea.state == 3}
      <td class="state-bg-green">🗸</td>
    {/if}
  </tr>
{/if}
