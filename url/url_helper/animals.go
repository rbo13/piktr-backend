package url_helper

// Animals returns a slice of
// animals.
func Animals() []string {
	return []string{"aardvark", "aardwolf", "abalone", "Abyssinian cat", "Abyssiniangroundhornbill", "acacia rat", "acornbarnacle", "acornweevil", "acornwoodpecker", "acouchi", "addax", "adder", "adelie penguin", "admiral", "admiralbutterfly", "adouri", "Africanaugurbuzzard", "AfricanBushviper", "Africancivet", "Africanclawedfrog", "AfricanElephant", "AfricanFisheagle", "AfricanGoldenCat", "AfricanGroundhornbill", "AfricanharrierHawk", "Africanhornbill", "AfricanJacana", "AfricanmoleSnake", "AfricanParadiseflycatcher", "Africanpiedkingfisher", "AfricanPorcupine", "African Rock Python", "African wild cat", "African wild dog", "agama", "agouta", "agouti", "Airedale", "aisan pied starling", "Akita inu", "Alabama map turtle", "Alaska jingle", "Alaskan husky", "Alaskan malamute", "albacore tuna", "albatross", "albertosaurus", "albino", "aldabra tortoise", "alligator", "alligator gar", "alligator snapping turtle", "allosaurus", "alpaca", "Alpine Black swallowtail butterfly", "Alpine goat", "Alpine road guide tiger beetle", "Altiplano chinchilla mouse", "Amazon dolphin", "Amazon parrot", "Amazon tree boa", "Amber pen shell", "American alligator", "American avocet", "American badger", "American bittern", "American Black vulture", "American cicada", "American crayfish", "American crocodile", "American crow", "American goldfinch", "American kestrel", "American lobster", "American marten", "American ratsnake", "American Red squirrel", "American river otter", "American Robin", "American toad", "American wigeon", "amethyst gem clam", "amethyst sunbird", "amethystine Python", "ammonite", "amoeba", "amphibian", "amphiuma", "Amur minnow", "Amur ratsnake", "Amur starfish", "anaconda", "anchovy", "Andean cat", "Andean cock-of-the-Rock", "Andean condor", "anemone", "anemone crab", "anemone shrimp", "Angel wing mussel", "angelfish", "anglerfish", "Angora", "angwantibo", "anhinga", "animated", "ankole", "ankole-Watusi", "annas hummingbird", "annelida", "annelid", "anole", "anopheles", "Antarctic fur seal", "Antarctic giant petrel", "anteater", "antelope", "antelope ground squirrel", "antipodes Green parakeet", "ant", "ant bear", "ant lion", "anura", "aoudad", "apatosaur", "ape", "aphid", "apis dorsata laboriosa", "aplomado falcon", "aquatic leech", "Arabian horse", "Arabian oryx", "Arabian wild cat", "aracari", "arachnid", "arawana", "archaeocete", "archaeopteryx", "archer fish", "Arctic Fox", "Arctic hare", "Arctic Wolf", "argali", "Argentine horned frog", "Argentine ruddy duck", "Argus fish", "Ariel toucan", "Arizona alligator lizard", "Ark shell", "armadillo", "armed crab", "armed nylon shrimp", "army ant", "arrow crab", "arrow worm", "arrowana", "arthropods", "aruanas", "Asian Constable butterfly", "Asian damselfly", "Asian elephant", "Asian lion", "Asian porcupine", "Asian Small-clawed otter", "Asian trumpetfish", "Asian water Buffalo", "Asiatic greater freshwater clam", "Asiatic lesser freshwater clam", "Asiatic mouflon", "Asiatic wild ass", "asp", "ass", "assassin bug", "Astarte", "astrangia coral", "Atlantic Black goby", "Atlantic blue tang", "Atlantic ridley turtle", "Atlantic sharpnose puffer", "Atlantic spadefish", "Atlas moth", "auk", "auklet", "aurochs", "Australian curlew", "Australian freshwater crocodile", "Australian fur seal", "Australian kestrel", "Australian shelduck", "Australian silky terrier", "avocet", "axis deer", "axolotl", "aye-aye", "Aztec ant", "azure-winged magpie", "azure vase", "babirusa", "baboon", "bactrian", "badger", "bagworm", "baiji", "bald eagle", "baleen whale", "balloonfish", "bandicoot", "banteng", "barasinga", "barasingha", "barb", "barbet", "barnacle", "barracuda", "basilisk", "Bass", "basset hound", "bat", "beagle", "bear", "bearded dragon", "beaver", "bee", "beetle", "Bell frog", "beluga whale", "bettong", "big-horned sheep", "bighorn", "bilby", "binturong", "Bird", "Bird of Paradise", "bison", "bittern", "Black bear", "Black fly", "Black panther", "Black rhino", "Black widow spider", "blackbird", "blackbuck", "blesbok", "blowfish", "blue Jay", "blue whale", "bluebird", "bluebottle", "bluefish", "boa", "boa constrictor", "boar", "bobcat", "bobolink", "bobwhite", "bongo", "booby", "boto", "boubou", "boutu", "bovine", "Brahman bull", "Brahman cow", "Brant", "bream", "brocket deer", "bronco", "brontosaurus", "Brown bear", "bubblefish", "Buck", "budgie", "bufeo", "Buffalo", "bufflehead", "bug", "bull mastiff", "bullfrog", "bumblebee", "bunny", "bunting", "burro", "Bush baby", "Bush squeaker", "bustard", "butterfly", "buzzard", "caecilian", "caiman", "caiman lizard", "calf", "Camel", "canary", "canine", "canvasback", "cape ghost frog", "capybara", "caracal", "cardinal", "caribou", "carp", "cassowary", "cat", "catbird", "Caterpillar", "catfish", "cattle", "caudata", "cavy", "centipede", "chafer", "chameleon", "chamois", "cheetah", "chevrotain", "chick", "chickadee", "chicken", "Chihuahua", "chimney Swift", "chimpanzee", "chinchilla", "Chinese crocodile lizard", "chipmunk", "chital", "chrysomelid", "chuckwalla", "chupacabra", "cicada", "cirriped", "civet", "clam", "cleaner wrasse", "clown anemone fish", "clumber", "coati", "cob", "cobra", "cockatiel", "cockatoo", "cocker spaniel", "cockroach", "cod", "coelacanth", "collard lizard", "Colt", "comet", "conch", "condor", "coney", "conure", "coot", "cooter", "copepod", "copperhead", "coqui", "coral", "cormorant", "corn Snake", "corydoras catfish", "cottonmouth", "cottontail", "cougar", "cow", "cowbird", "cowrie", "coyote", "coypu", "crab", "Crane", "crayfish", "cricket", "crocodile", "crocodile skink", "crossbill", "crow", "crown-of-thorns starfish", "crustacean", "cub", "cuckoo", "curassow", "curlew", "cuscus", "cusimanse", "cuttlefish", "dog", "dolphin", "duck", "bald eagle", "Golden eagle", "egret", "eidolon", "common eland", "giant eland", "African elephant", "Asian elephant", "scrawled filefish", "peregrine falcon", "carribean flamingo", "Chilean flamingo", "lesser flamingo", "Malaysian flying Fox", "Arctic Fox", "fennec Fox", "Cuban tree Fox", "South American ornate horned Fox", "tawny frogmouth", "frog", "mantella frog", "poison arrow frog", "purple gallinule", "gars", "addra gazelle", "Dorcas gazelle", "slender-horned gazelle", "leopard gecko", "Solomon island gecko", "Tokay gecko", "gharial", "White-cheeked Gibbon", "gibbons", "Gila monster", "giraffe", "domestic goat", "Abyssinian blue-winged goose", "African Pygmy goose", "bar-headed goose", "cape barren goose", "magpie goose", "Western lowland gorilla", "groupers", "bluestriped grunt", "French grunt", "helmeted guineafowl", "vulturine guineafowl", "gulls", "hackee", "haddock", "hadrosaurus", "hag-fish", "hairstreak", "hake", "halcyon", "halibut", "halicore", "hamadryad", "hamadryas", "hammerhead Bird", "hammerhead shark", "hammerkop", "hamster", "hanuman-monkey", "hapuka", "hapuku", "harbor-porpoise", "harbor-seal", "hare", "harp-seal", "harpy-eagle", "harrier", "harrier hawk", "Hart", "hartebeest", "harvest mouse", "harvestmen", "hatchet-fish", "Hawaiian-Monk seal", "hawk", "hedgehog", "heifer", "hellbender", "hen", "herald", "Hercules-beetle", "hermit crab", "heron", "Herring", "heterodontosaurus", "hind", "hippo", "hippopotamus", "hoatzin", "hog", "hogget", "hoiho", "hoki", "homalocephale", "honey-badger", "honeybee", "honey-creeper", "honeyeater", "hoopoe", "Horn-shark", "horned-toad", "horned-viper", "hornbill", "hornet", "horse", "horsefly", "horseshoe bat", "horseshoe crab", "hound", "House-mouse", "housefly", "howler-monkey", "huemul (deer)", "huia", "hummingbird", "humpback whale", "husky", "hydatid-tapeworm", "Hydra", "hyena", "hylaeosaurus", "hypacrosaurus", "hypsilophodon", "hyracotherium", "hyrax", "sacred ibis", "scarlet ibis", "ibises", "Cuban iguana", "Green iguana", "rhinoceros iguana", "impala", "jabiru", "jackal", "jackrabbit", "jaeger", "Jaguar", "jaguarundi", "janenschia", "Jay", "jellyfish", "Jenny", "jerboa", "Joey", "John dory", "junco", "June bug", "Western Gray kangaroo", "White-collared kingfisher", "koala", "Ugandan kob", "komodo dragon", "laughing kookaburra", "greater kudu", "leafy sea dragon", "Black & White ruffed lemur", "Red ruffed lemur", "ring-tailed lemur", "lemurs", "leopard", "African lion", "lionfish - scorpionfish - stonefish", "Mexican beaded lizard", "lizards", "llama", "lookdown", "blue-streaked lory", "chattering lory", "dusky lory", "Green-naped lory", "southeastern lubber", "blue & gold macaw", "blue-throated macaw", "Green-winged macaw", "hyacinth macaw", "Mexican military macaw", "macaws", "manatees", "common marmoset", "slender-tailed meerkat", "hooded merganser", "Honduran milksnake", "crocodile monitor", "Malayan water monitor", "White-throated monitor", "common moorhen", "morays", "Indian muntjac", "Golden-crested mynah", "nabarlek", "nag", "naga", "nagapies", "naked mole rat", "nandine", "nandoo", "nandu", "narwhal", "narwhale", "natterjack toad", "nauplius", "Nautilus", "needle fish", "needletail", "nematode", "nene", "neon blue guppy", "neon blue hermit crab", "neon dwarf gourami", "neon rainbow fish", "neon Red guppy", "neon tetra", "nerka", "nettlefish", "Newfoundland dog", "newt", "newt nutria", "night crawler", "night heron", "nighthawk", "Nightingale", "nightjar", "nilgai", "nine banded armadillo", "noctilio", "noctule", "noddy", "noolbenger", "northern cardinals", "northern elephant seal", "northern flying squirrel", "northern fur seal", "northern hairy-nosed wombat", "northern Pike", "northern sea horse", "northern spotted owl", "Norway lobster", "Norway rat", "Nubian goat", "nudibranch", "numbat", "nurse shark", "nutcracker", "nuthatch", "nutria", "nyala", "nymph", "Virginia opossum", "bornean orangutan", "scimitar-horned oryx", "ostrich", "Asian Small-clawed otter", "Congo clawless otter", "giant river otter", "sea otter", "barn owl", "great horned owl", "Florida panther", "derbyan parakeet", "African Gray parrot", "Brazilian hawk-headed parrot", "grand eclectus parrot", "parrotfish", "Queen parrotfish", "parrots", "Indian peafowl", "peccaries", "Brown pelican", "Adélie penguin", "chinstrap penguin", "gentoo penguin", "penguins", "Golden pheasant", "pot-bellied pig", "lesser Bahama pintail", "piranhas", "Atlantic porkfish", "primates", "pufferfish", "porcupinefish", "tufted puffin", "puffins", "Burmese Python", "carpet Python", "Green tree Python", "Royal Python", "quagga - extinct zebra", "quahog - type of clam", "quail - type of ground Bird", "Queen Snake", "Queen ant", "Queen bee", "quetzal - colorful Bird of S. America", "quokka", "quoll", "old world rabbit", "raccoons", "light-footed clapper rail", "canebrake rattlesnake", "dusky Pygmy rattlesnake", "Eastern diamondback rattlesnake", "spotted eagle Ray", "rays", "electric rays", "Black rhinoceros", "White rhinoceros", "tiger salamander", "sawfishes", "emperor scorpion", "California sea lion", "seahorses", "crabeater seal", "harbor seal", "hooded seal", "leopard seal", "Ross seal", "Weddell seal", "Monk seals", "sergeant Major", "serval", "bonnethead shark", "epaulette shark", "great White shark", "leopard shark", "nurse shark", "Pacific blacktip reef shark", "whale shark", "sharks", "Barbary sheep", "domestic sheep", "Australian shelduck", "cape shelduck", "Argentine Red shoveler", "Australian shoveler", "common shoveler", "Eastern blue-tongued skink", "Black rat Snake", "Eastern corn Snake", "Eastern indigo Snake", "Eastern King Snake", "Florida cottonmouth Snake", "Florida King Snake", "Florida pine Snake", "Gray rat Snake", "scarlet King Snake", "yellow rat Snake", "snakes", "sea snakes", "Atlantic spadefish", "spiders", "Golden-breasted starling", "Black-necked stilt", "marabou stork", "sugar glider", "sunbittern", "Black swan", "Black-necked swan", "coscoroba swan", "Beryl-spangled tanager", "blue & Black tanager", "blue tang", "tapirs", "Western tarsier", "American Green wing teal", "Baikal teal", "cape teal", "cinnamon teal", "garganey teal", "Hottentot teal", "marbled teal", "Madagascar tenrec", "Bengal tiger", "Marine toad", "Oriental fire-bellied toad", "tomistoma", "African spurred tortoise", "aldabra tortoise", "yellow-footed tortoise", "keel-billed toucan", "toco toucan", "Gray-winged trumpeter", "Grey turaco", "Guinea turaco", "Red-crested turaco", "White-cheeked turaco", "sea turtles", "uakari", "Uganda kob", "uinta ground squirrel", "umbrella Bird", "umbrette", "unau", "ungulate", "unicorn", "upupa", "urchin", "urial", "uromastyx maliensis", "uromastyx spinipes", "urson", "urubu", "urus", "urutu", "urva", "Utah prairie dog", "vaquita", "Black vulture", "King vulture", "lappet-faced vulture", "vultures", "old world vultures", "vampire bat", "vaquita", "veery", "velociraptor", "velvet crab", "velvet worm", "venomous Snake", "verdin", "vervet", "vicuna", "viper", "viper squid", "viperfish", "vireo", "Virginia opossum", "vixen", "vole", "volvox", "vulpes velox", "vulpes vulpes", "walrus", "warthog", "defassa waterbuck", "beluga whale", "blue whale", "false killer whale", "Gray whale", "killer whale", "Short-finned pilot whale", "sperm whale", "whales", "endangered whales", "American wigeon", "chiloe wigeon", "Eastern White-bearded (gnu) wildebeest", "Gray Wolf", "maned Wolf", "X-Ray fish", "X-Ray tetra", "xanclomys", "xanthareel", "xantus", "xantus murrelet", "xeme", "xenarthra", "xenoposeidon", "xenops", "xenopterygii", "xenopus", "xenotarsosaurus", "xenurine", "xenurus unicinctus", "xerus", "xiaosaurus", "xinjiangovenator", "xiphias", "xiphias gladius", "xiphosuran", "xoloitzcuintli", "xoni", "xuanhanosaurus", "xuanhuaceratops", "xuanhuasaurus", "yaffle", "yak", "yapok", "yard ant", "yearling", "yellow bellied marmot", "yellow belly lizard", "yellowhammer", "yellowjacket", "yellowlegs", "yellowthroat", "yeti", "ynambu", "Yorkshire terrier", "Yosemite toad", "yucker", "zander", "Zanzibar Day gecko", "zebra", "zebra dove", "zebra Finch", "zebra-tailed lizard", "zebrafish", "zebu", "zenaida", "zeren", "zethus spinipes", "zethus Wasp", "zigzag salamander", "zone-tailed pigeon", "zooplankton", "zopilote", "zorilla"}
}
