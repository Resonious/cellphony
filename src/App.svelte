<script>
  export let name;

  const lambda = 'https://uueg3x70bg.execute-api.us-west-2.amazonaws.com/default/cellphonyRtcNotify';

  let connection;
  let channel;
  let myID = Math.random().toString(36).substr(2, 9);
  let myOffer;

  let pusher = new Pusher('151c09f47c37e803cb69', {
    cluster: 'ap3'
  });

  let remoteOffer = (location.hash.length > 0) &&
    JSON.parse(atob(location.hash.substr(1)).toString());

  $: remoteID = (!!remoteOffer) &&
    remoteOffer.id;

  let currentHost = location.toString().split('#')[0];

  $: shareUrl = (!!myOffer) &&
    `${currentHost}#${btoa(JSON.stringify({ id: myID, offer: myOffer }))}`;


  async function host() {
    connection = new RTCPeerConnection();
    channel = connection.createDataChannel('messageChannel');
    channel.addEventListener('open', onChannelStatusChange);
    channel.addEventListener('close', onChannelStatusChange);

    pusher.subscribe(myID).bind('join', data => {
      console.log('SOMEBODY WANTS TO JOIN');
      console.log(data);
    });

    let offer = await connection.createOffer();
    await connection.setLocalDescription(offer);
    console.log(offer);
    myOffer = offer;
    window.myOffer = offer;
  }

  async function join() {
    connection = new RTCPeerConnection();
    connection.addEventListener('datachannel', ch => {
      channel = ch;
      channel.addEventListener('open', onChannelStatusChange);
      channel.addEventListener('close', onChannelStatusChange);
      console.log('I GOT A CHANNEL !!!');
    });

    await fetch(lambda, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      redirect: 'follow',
      body: JSON.stringify({ id: myID, room: remoteID, answer: JSON.stringify(answer) })
    });

    console.log('sent...');
  }

  // So I know what's going ON
  function onChannelStatusChange(event) {
    console.log(`Channel state: ${channel.state}`);
  }
</script>

<main>
  <h1>Hello {name}! I am {myID}</h1>
  {#if remoteID}
    <h2>Joining {remoteID}</h2>
  {:else if !myOffer}
    <button on:click={host}>Host</button>
  {/if}

  {#if remoteID}
    <button on:click={join}>Join {remoteID}</button>
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
