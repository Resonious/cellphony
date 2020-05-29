<script>
  export let name;

  let connection;
  let channel;
  let myOffer;

  let remoteOffer = location.hash;

  let currentHost = location.toString().split('#')[0];

  $: shareUrl = !!myOffer && `${currentHost}#${btoa(JSON.stringify(myOffer))}`;

  async function host() {
    connection = new RTCPeerConnection();
    channel = connection.createDataChannel('messageChannel');
    channel.addEventListener('open', onChannelStatusChange);
    channel.addEventListener('close', onChannelStatusChange);

    let offer = await connection.createOffer();
    await connection.setLocalDescription(offer);
    console.log(offer);
    myOffer = offer;
  }

  function join() {
    connection = new RTCPeerConnection();
    connection.addEventListener('datachannel', ch => {
      channel = ch;
      channel.addEventListener('open', onChannelStatusChange);
      channel.addEventListener('close', onChannelStatusChange);
      console.log('I GOT A CHANNEL !!!');
    });
  }

  // So I know what's going ON
  function onChannelStatusChange(event) {
    console.log(`Channel state: ${channel.state}`);
  }
</script>

<main>
  <h1>Hello {name}!</h1>

  <button on:click={host}>Host</button>
  {#if remoteOffer && remoteOffer.length > 0}
    <button on:click={join}>Join {remoteOffer}</button>
  {/if}

  {#if shareUrl}
    <p><a href='{shareUrl}'>Share link</a></p>
  {/if}
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
