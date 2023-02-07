<script>
  export let idea;
  let firstLink = "https://google.com";
  //$: let firstLink = idea === null ? "":idea.manual_url;
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
      <div />
      <div />
    </td>
    <td />
    <td />
  </tr>
{/if}
