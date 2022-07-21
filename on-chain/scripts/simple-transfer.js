
const { ethers, upgrades } = require("hardhat");

const { MaxUint256 } = require("ethers");

const overrides = {
  gasLimit: 9999999
}

async function main() {

  const [k1,deployer,b,c, a] = await ethers.getSigners();

  console.log("sender: ", a.address);

 const txData = {
    from: a.address,
    to: "0xC0F0Ce81234776f7bf30e479fB0Ac518034818B6",
    value: ethers.utils.parseEther("0.005"),
    // nonce: await provider.getTransactionCount(deployer.address, "latest"),
    // gasLimit: ethers.utils.hexlify(9999999),
    // gasPrice: ethers.utils.hexlify(25277782739),
    data: ethers.utils.hexlify(ethers.utils.toUtf8Bytes(`Dory and Meta: A Tale of Metamask
    Chapter #1: An interview
    “Morning, boss! Here’s the summary report for our transactions last week.”
    In a simple but nicely-decorated office, a pretty female fox secretary put a file folder onto his supervisor’s table.
    ‘Thanks Kim. And, again, just call me Meta.’ The supervisor, named Meta, was a young male fox in his twenties. As the head of the entire Metamask Department, Meta had elegantly-maintained orange fur and he’s dressed in a casual smart suit. Handsome, talented and wealthy, Meta was very popular within the entire Enterprise of ConsenSys, the mother company of the Metamask Department. Meta’s job was significant to the company, as he’s here to help VIP clients of his department with their foreign exchange transactions.
    ‘So, Opensea is still #1 in gas burn, and the foreign exchange business is still thriving.’ Meta quickly browsed the summary report and nodded, happy. ‘This planet is getting busier and more prosperous. I was wondering how we look now compared to the BTC planet?’
    ‘They are doing just as great.’ Kim smiled. ‘As you know, the prosperity of their planet and ours tend to rise and fall on a similar pattern.’
    ‘True. But clients don’t seem to be traveling as often to their planet as to this one.’ Meta leaned over to his window on the 60th floor, and enjoyed the birdview of the entire enterprise campus. As he looked further away, he saw the sunrise coming right on the horizon. ‘See that? How beautiful! I just love the scenery of this ETH planet, which is why I chose to stay here most of the time.’
    ‘Boss, I think you need to take a rest.’ Kim hesitated a bit but decided to speak up. ‘I always consider myself an early riser, but everyday when I got here, you’ve been in the office working.’
    
     ‘I’d like to get ahead of the sunrise. That’s my favorite part of the day. And you know just how busy we have been recently, don’t you?’ Meta took a long breath and smiled. ‘Speaking of that, next month I’ll need to visit the Polygon and BSC planets as our business has been thriving there recently. So I need you to help me settle everything here before I go.’
    ‘Don’t worry, we are doing very well here.’ Kim said confidently. ‘The only hindrance was that not everyone out there understands what exactly we have been doing and how important our work is. If they do, our business would be expanding by ten fold or even more.’
    ‘Bingo! I’ve got that part planned.’ Meta made a finger click. ‘In a few moments, a journalist will be visiting us and giving me an interview. With her help, I was hoping we can get better known – especially by those people who live on the planets in outer space.’
    ‘Oh, that’s the reason for the interview today? I’ve been wondering since when you had time for an interview from a nobody.’ Kim quickly checked Meta’s schedule. ‘Is it someone from... hmm... StoryDAO? What’s that?’
    ‘Some DAO, apparently. DAOs are all over the place nowadays.’ Meta lay back in his chair, putting his right leg on his left. ‘Oh, you don’t know what a DAO is? It’s called decentralized autonomous organization. Consider that as a country with complete democracy. They let all of their civilians vote whenever they are about to make a big decision of some sort.’
    ‘Sounds interesting. So the journalist came from a country of that style, and her name is...’
    Right then, someone knocked at the door and interrupted Kim. She opened the door, and saw a little monkey outside.
    ‘Morning! My name is Dory. I had an appointment with Mr. Mask.’ The little monkey sounded like a girl. ‘Have I come to the right place?’
    ‘That’s me. Come on in please, miss.’ Meta stood up as a courtesy, and signaled Kim to leave. ‘It must have been a long trip for you. I hope it’s been smooth.’
    ‘It is. This campus is so big and hard to miss.’ Dory looked quite excited at the same time a bit nervous. ‘It’s such a pleasure to meet you, sir. To be honest, I didn’t expect that someone like you would have time for my interview!’
    
     ‘Me neither... cough cough!’ Meta hurriedly mumbled as he accidentally slipped it out. ‘Well, to be completely honest with you, we are indeed so very busy these days. But I thought it’s a good idea to switch my mind once in a while and make some new friends.’
    ‘What an honor!’ Dory took out a pen and a booklet. ‘I fear to waste your time then. Shall we begin our interview already?’
    ‘Sure. What would you like to know?’ Meta took a sip of coffee.
    ‘As far as I understand, the main business of Metamask is to help clients complete their financial transactions, following some universal trading protocols.’ Dory carefully worded her speech. ‘And your mother company, ConsenSys, has contributed to the constitution process of many trading protocols.’
    ‘Correct. Those trading protocols are significant to our world, and they’ve been constantly maintained in the respective departments.’ Meta pointed to his window, through which a number of tall buildings could be seen. Dory saw that on the top of each building there was a big sign, marking the names of the departments. The tallest building has a sign of “ERC-20” on it.
    ‘I do know that ERC-20 is one of the most important protocols worldwide. It’s almost like a financial Law.’ Dory couldn’t help marveling, as it’s her first time to see this legendary building. ‘And ERC-20 is so popular that it’s not only in effect on this planet, but also has derivations on many others such as Polygon and Avalanche. What an awesome job you guys have done!’
    ‘Well... we did do our part, but we are not the creator of ERC-20. It was created by the Ethereum genesis team, led by Mr. Vitalik Buterin.’ Meta said in a tone of respect. ‘But we are certainly proud to be of help, not only to ERC-20, but also to many other important ones, such as ERC-721 and ERC-1155.’
    ‘The ones for NFTs?’ Dory’s eyes shone up. ‘I’m a big fan of those. We have many NFTs in our DAO.’
    ‘Non-fungible tokens, correct. But don’t you dare to advertise to me about your NFTs, or our security guys will fine you for penalties.’ Meta teased her. ‘In case you don’t know, NFTs always give us headaches. For many times when a NFT collection was released, so many people got on their way to Opensea for purchase. That caused severe traffic jams and caused the gas fee to become scarily high.’
    ‘Just to clarify, that had nothing to do with us...’ Dory said carefully. ‘Although we do plan to make an NFT collection from this interview today...’
    
     ‘You plan to do that?’ Meta seemed interested. ‘Don’t take me wrong, although it’s a headache sometimes, we do like NFTs as they have become a big part of our business nowadays. As you know, high gas fee means that we are busy, which makes us a busy-ness.’
    ‘Well, I hope we can be of help then.’ Dory made a grimace. ‘If you sir could show me some fun parts about your business, I’m going to work something out, starting from a story – you know, we have plenty of authors in our DAO.’
    ‘You do? It sounds like a plan then.’ Meta walked to the door, smiling. ‘I can show you around, but I’m not sure if our daily work looks fun to you. Although we carry out important tasks, it’s quite tedious if you do the same thing over and over again.’
    ‘Try me!’ Dory shot up from her chair, excited. She followed Meta out of his office to tour the entire building of the Metamask department. Kim had just brought coffee to the office, and she’s surprised to see the two leaving it.
    ‘You come at the right moment, Kim. Bring a client to the VIP room, and I’ll be there soon to show our monkey friend how we handle our business everyday.’ Meta hastened Kim away to get ready, meanwhile he showed Dory around the building. ‘As you know, we are fortunate to be living in a planet of fairness, where everyone could gain his income based on the proof of his work, and we call this the PoW spirit. Most of the planets believe in the PoW spirit, starting from the very first one where our civilization started.’
    ‘You mean the BTC planet?’ Dory asked. ‘I’ve never been there and I always wanted to.’
    ‘I don’t go there either. It’s a planet with rich mineral resources, but the natural environment isn’t as friendly as here, so people didn’t build many facilities on it.’ Meta agreed. ‘I was told that only miners stay on that planet, and most of them possess great fortune.’
    As the two chatted along their way towards the VIP room, Dory saw many foxes running back and forth with suitcases of different sizes. Seeing her curious face, Meta explained. ‘Those suitcases contain the assets of our clients. Most of the time they are tokens, but sometimes it could be NFTs in there. Our folks are getting ready to escort those assets to an exchange for trading.’
    ‘Which exchange are they going to?’ Dory asked.
    ‘It really depends on the clients.’ Meta patiently explained. ‘As you may already know, there are many different exchanges, some are centralized, short as CEXs, meaning there’s a verified entity running it. But those have nothing to do with us because we only
    
     work with DEXs. That is decentralized exchanges, meaning that the entity behind it is anonymous. They function in a similar way though.’
    ‘There are two major DEXs on our planet of ETH, which are Uniswap and Sushiswap.’ Seeing Dory nodded, Meta continued. ‘Let’s make a bet - which one do you guess we are going to this time?’
    ‘I like Sushi. I’ll bet on it.’ Dory’s eyes twinkled. ‘Hmm... I thought monkeys like banana the best, but sadly there’s no bananaswap.’ Meta teased as they proceeded to the VIP room. ‘I’ll go for Uni then. Let’s see which client Kim will lead to us and what he wants.’
    (to be continued)`))
};

console.log("start tx.")

const tx =await a.sendTransaction(txData, overrides);

await tx.wait();

console.log("complete tx.")
//   console.log("complete mint.")
}


main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
  