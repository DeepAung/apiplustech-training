// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Mutation struct{}

type Pokemon struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Category    string           `json:"category"`
	Types       []PokemonType    `json:"types"`
	Abilities   []PokemonAbility `json:"abilities"`
}

type PokemonInput struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Category    string           `json:"category"`
	Types       []PokemonType    `json:"types"`
	Abilities   []PokemonAbility `json:"abilities"`
}

type Query struct{}

type PokemonAbility string

const (
	PokemonAbilityAdaptability     PokemonAbility = "ADAPTABILITY"
	PokemonAbilityAftermath        PokemonAbility = "AFTERMATH"
	PokemonAbilityAirLock          PokemonAbility = "AIR_LOCK"
	PokemonAbilityAngerPoint       PokemonAbility = "ANGER_POINT"
	PokemonAbilityAngerShell       PokemonAbility = "ANGER_SHELL"
	PokemonAbilityAnticipation     PokemonAbility = "ANTICIPATION"
	PokemonAbilityArenaTrap        PokemonAbility = "ARENA_TRAP"
	PokemonAbilityArmorTail        PokemonAbility = "ARMOR_TAIL"
	PokemonAbilityAromaVeil        PokemonAbility = "AROMA_VEIL"
	PokemonAbilityAuraBreak        PokemonAbility = "AURA_BREAK"
	PokemonAbilityBadDreams        PokemonAbility = "BAD_DREAMS"
	PokemonAbilityBallFetch        PokemonAbility = "BALL_FETCH"
	PokemonAbilityBattery          PokemonAbility = "BATTERY"
	PokemonAbilityBattleArmor      PokemonAbility = "BATTLE_ARMOR"
	PokemonAbilityBeadsOfRuin      PokemonAbility = "BEADS_OF_RUIN"
	PokemonAbilityBeastBoost       PokemonAbility = "BEAST_BOOST"
	PokemonAbilityBerserk          PokemonAbility = "BERSERK"
	PokemonAbilityBigPecks         PokemonAbility = "BIG_PECKS"
	PokemonAbilityBlaze            PokemonAbility = "BLAZE"
	PokemonAbilityBulletproof      PokemonAbility = "BULLETPROOF"
	PokemonAbilityCheekPouch       PokemonAbility = "CHEEK_POUCH"
	PokemonAbilityChillingNeigh    PokemonAbility = "CHILLING_NEIGH"
	PokemonAbilityChlorophyll      PokemonAbility = "CHLOROPHYLL"
	PokemonAbilityClearBody        PokemonAbility = "CLEAR_BODY"
	PokemonAbilityCloudNine        PokemonAbility = "CLOUD_NINE"
	PokemonAbilityColorChange      PokemonAbility = "COLOR_CHANGE"
	PokemonAbilityComatose         PokemonAbility = "COMATOSE"
	PokemonAbilityCommander        PokemonAbility = "COMMANDER"
	PokemonAbilityCompetitive      PokemonAbility = "COMPETITIVE"
	PokemonAbilityCompoundEyes     PokemonAbility = "COMPOUND_EYES"
	PokemonAbilityContrary         PokemonAbility = "CONTRARY"
	PokemonAbilityCorrosion        PokemonAbility = "CORROSION"
	PokemonAbilityCottonDown       PokemonAbility = "COTTON_DOWN"
	PokemonAbilityCudChew          PokemonAbility = "CUD_CHEW"
	PokemonAbilityCursedBody       PokemonAbility = "CURSED_BODY"
	PokemonAbilityCuteCharm        PokemonAbility = "CUTE_CHARM"
	PokemonAbilityDamp             PokemonAbility = "DAMP"
	PokemonAbilityDancer           PokemonAbility = "DANCER"
	PokemonAbilityDark             PokemonAbility = "DARK"
	PokemonAbilityAura             PokemonAbility = "AURA"
	PokemonAbilityDauntlessShield  PokemonAbility = "DAUNTLESS_SHIELD"
	PokemonAbilityDazzling         PokemonAbility = "DAZZLING"
	PokemonAbilityDefeatist        PokemonAbility = "DEFEATIST"
	PokemonAbilityDefiant          PokemonAbility = "DEFIANT"
	PokemonAbilityDisguise         PokemonAbility = "DISGUISE"
	PokemonAbilityDownload         PokemonAbility = "DOWNLOAD"
	PokemonAbilityDragonsMaw       PokemonAbility = "DRAGONS_MAW"
	PokemonAbilityDrizzle          PokemonAbility = "DRIZZLE"
	PokemonAbilityDrought          PokemonAbility = "DROUGHT"
	PokemonAbilityDrySkin          PokemonAbility = "DRY_SKIN"
	PokemonAbilityEarlyBird        PokemonAbility = "EARLY_BIRD"
	PokemonAbilityEarthEater       PokemonAbility = "EARTH_EATER"
	PokemonAbilityEffectSpore      PokemonAbility = "EFFECT_SPORE"
	PokemonAbilityElectricSurge    PokemonAbility = "ELECTRIC_SURGE"
	PokemonAbilityElectromorphosis PokemonAbility = "ELECTROMORPHOSIS"
	PokemonAbilityEmergencyExit    PokemonAbility = "EMERGENCY_EXIT"
	PokemonAbilityFairyAura        PokemonAbility = "FAIRY_AURA"
	PokemonAbilityFilter           PokemonAbility = "FILTER"
	PokemonAbilityFlameBody        PokemonAbility = "FLAME_BODY"
	PokemonAbilityFlashFire        PokemonAbility = "FLASH_FIRE"
	PokemonAbilityFlowerGift       PokemonAbility = "FLOWER_GIFT"
	PokemonAbilityFlowerVeil       PokemonAbility = "FLOWER_VEIL"
	PokemonAbilityFluffy           PokemonAbility = "FLUFFY"
	PokemonAbilityForecast         PokemonAbility = "FORECAST"
	PokemonAbilityForewarn         PokemonAbility = "FOREWARN"
	PokemonAbilityFriendGuard      PokemonAbility = "FRIEND_GUARD"
	PokemonAbilityFrisk            PokemonAbility = "FRISK"
	PokemonAbilityFullMetalBody    PokemonAbility = "FULL_METAL_BODY"
	PokemonAbilityFurCoat          PokemonAbility = "FUR_COAT"
	PokemonAbilityGluttony         PokemonAbility = "GLUTTONY"
	PokemonAbilityGoodAsGold       PokemonAbility = "GOOD_AS_GOLD"
	PokemonAbilityGooey            PokemonAbility = "GOOEY"
	PokemonAbilityGrassySurge      PokemonAbility = "GRASSY_SURGE"
	PokemonAbilityGrimNeigh        PokemonAbility = "GRIM_NEIGH"
	PokemonAbilityGuardDog         PokemonAbility = "GUARD_DOG"
	PokemonAbilityGulpMissile      PokemonAbility = "GULP_MISSILE"
	PokemonAbilityGuts             PokemonAbility = "GUTS"
	PokemonAbilityHadronEngine     PokemonAbility = "HADRON_ENGINE"
	PokemonAbilityHealer           PokemonAbility = "HEALER"
	PokemonAbilityHeatproof        PokemonAbility = "HEATPROOF"
	PokemonAbilityHeavyMetal       PokemonAbility = "HEAVY_METAL"
	PokemonAbilityHoneyGather      PokemonAbility = "HONEY_GATHER"
	PokemonAbilityHospitality      PokemonAbility = "HOSPITALITY"
	PokemonAbilityHugePower        PokemonAbility = "HUGE_POWER"
	PokemonAbilityHungerSwitch     PokemonAbility = "HUNGER_SWITCH"
	PokemonAbilityHustle           PokemonAbility = "HUSTLE"
	PokemonAbilityHydration        PokemonAbility = "HYDRATION"
	PokemonAbilityHyperCutter      PokemonAbility = "HYPER_CUTTER"
	PokemonAbilityIceBody          PokemonAbility = "ICE_BODY"
	PokemonAbilityIceFace          PokemonAbility = "ICE_FACE"
	PokemonAbilityIlluminate       PokemonAbility = "ILLUMINATE"
	PokemonAbilityIllusion         PokemonAbility = "ILLUSION"
	PokemonAbilityImmunity         PokemonAbility = "IMMUNITY"
	PokemonAbilityInfiltrator      PokemonAbility = "INFILTRATOR"
	PokemonAbilityInnardsOut       PokemonAbility = "INNARDS_OUT"
	PokemonAbilityInnerFocus       PokemonAbility = "INNER_FOCUS"
	PokemonAbilityInsomnia         PokemonAbility = "INSOMNIA"
	PokemonAbilityIntimidate       PokemonAbility = "INTIMIDATE"
	PokemonAbilityIntrepidSword    PokemonAbility = "INTREPID_SWORD"
	PokemonAbilityIronBarbs        PokemonAbility = "IRON_BARBS"
	PokemonAbilityIronFist         PokemonAbility = "IRON_FIST"
	PokemonAbilityJustified        PokemonAbility = "JUSTIFIED"
	PokemonAbilityKeenEye          PokemonAbility = "KEEN_EYE"
	PokemonAbilityKlutz            PokemonAbility = "KLUTZ"
	PokemonAbilityLeafGuard        PokemonAbility = "LEAF_GUARD"
	PokemonAbilityLevitate         PokemonAbility = "LEVITATE"
	PokemonAbilityLightMetal       PokemonAbility = "LIGHT_METAL"
	PokemonAbilityLightningRod     PokemonAbility = "LIGHTNING_ROD"
	PokemonAbilityLimber           PokemonAbility = "LIMBER"
	PokemonAbilityLingeringAroma   PokemonAbility = "LINGERING_AROMA"
	PokemonAbilityLiquidOoze       PokemonAbility = "LIQUID_OOZE"
	PokemonAbilityMagicGuard       PokemonAbility = "MAGIC_GUARD"
	PokemonAbilityMagician         PokemonAbility = "MAGICIAN"
	PokemonAbilityMagmaArmor       PokemonAbility = "MAGMA_ARMOR"
	PokemonAbilityMagnetPull       PokemonAbility = "MAGNET_PULL"
	PokemonAbilityMarvelScale      PokemonAbility = "MARVEL_SCALE"
	PokemonAbilityMegaLauncher     PokemonAbility = "MEGA_LAUNCHER"
	PokemonAbilityMerciless        PokemonAbility = "MERCILESS"
	PokemonAbilityMinus            PokemonAbility = "MINUS"
	PokemonAbilityMistySurge       PokemonAbility = "MISTY_SURGE"
	PokemonAbilityMoldBreaker      PokemonAbility = "MOLD_BREAKER"
	PokemonAbilityMotorDrive       PokemonAbility = "MOTOR_DRIVE"
	PokemonAbilityMoxie            PokemonAbility = "MOXIE"
	PokemonAbilityMultitype        PokemonAbility = "MULTITYPE"
	PokemonAbilityMummy            PokemonAbility = "MUMMY"
	PokemonAbilityMyceliumMight    PokemonAbility = "MYCELIUM_MIGHT"
	PokemonAbilityNaturalCure      PokemonAbility = "NATURAL_CURE"
	PokemonAbilityNeutralizingGas  PokemonAbility = "NEUTRALIZING_GAS"
	PokemonAbilityNoGuard          PokemonAbility = "NO_GUARD"
	PokemonAbilityNormalize        PokemonAbility = "NORMALIZE"
	PokemonAbilityOblivious        PokemonAbility = "OBLIVIOUS"
	PokemonAbilityOpportunist      PokemonAbility = "OPPORTUNIST"
	PokemonAbilityOrichalcumPulse  PokemonAbility = "ORICHALCUM_PULSE"
	PokemonAbilityOvercoat         PokemonAbility = "OVERCOAT"
	PokemonAbilityOvergrow         PokemonAbility = "OVERGROW"
	PokemonAbilityOwnTempo         PokemonAbility = "OWN_TEMPO"
	PokemonAbilityPickpocket       PokemonAbility = "PICKPOCKET"
	PokemonAbilityPickup           PokemonAbility = "PICKUP"
	PokemonAbilityPlus             PokemonAbility = "PLUS"
	PokemonAbilityPoisonHeal       PokemonAbility = "POISON_HEAL"
	PokemonAbilityPoisonPoint      PokemonAbility = "POISON_POINT"
	PokemonAbilityPoisonPuppeteer  PokemonAbility = "POISON_PUPPETEER"
	PokemonAbilityPoisonTouch      PokemonAbility = "POISON_TOUCH"
	PokemonAbilityPowerSpot        PokemonAbility = "POWER_SPOT"
	PokemonAbilityPrankster        PokemonAbility = "PRANKSTER"
	PokemonAbilityPressure         PokemonAbility = "PRESSURE"
	PokemonAbilityPrismArmor       PokemonAbility = "PRISM_ARMOR"
	PokemonAbilityProtosynthesis   PokemonAbility = "PROTOSYNTHESIS"
	PokemonAbilityPsychicSurge     PokemonAbility = "PSYCHIC_SURGE"
	PokemonAbilityPunkRock         PokemonAbility = "PUNK_ROCK"
	PokemonAbilityPurePower        PokemonAbility = "PURE_POWER"
	PokemonAbilityPurifyingSalt    PokemonAbility = "PURIFYING_SALT"
	PokemonAbilityQuarkDrive       PokemonAbility = "QUARK_DRIVE"
	PokemonAbilityQueenlyMajesty   PokemonAbility = "QUEENLY_MAJESTY"
	PokemonAbilityQuickFeet        PokemonAbility = "QUICK_FEET"
	PokemonAbilityRksSystem        PokemonAbility = "RKS_SYSTEM"
	PokemonAbilityRainDish         PokemonAbility = "RAIN_DISH"
	PokemonAbilityRattled          PokemonAbility = "RATTLED"
	PokemonAbilityReceiver         PokemonAbility = "RECEIVER"
	PokemonAbilityReckless         PokemonAbility = "RECKLESS"
	PokemonAbilityRefrigerate      PokemonAbility = "REFRIGERATE"
	PokemonAbilityRegenerator      PokemonAbility = "REGENERATOR"
	PokemonAbilityRipen            PokemonAbility = "RIPEN"
	PokemonAbilityRivalry          PokemonAbility = "RIVALRY"
	PokemonAbilityRockHead         PokemonAbility = "ROCK_HEAD"
	PokemonAbilityRoughSkin        PokemonAbility = "ROUGH_SKIN"
	PokemonAbilityRunAway          PokemonAbility = "RUN_AWAY"
	PokemonAbilitySandForce        PokemonAbility = "SAND_FORCE"
	PokemonAbilitySandRush         PokemonAbility = "SAND_RUSH"
	PokemonAbilitySandSpit         PokemonAbility = "SAND_SPIT"
	PokemonAbilitySandStream       PokemonAbility = "SAND_STREAM"
	PokemonAbilitySandVeil         PokemonAbility = "SAND_VEIL"
	PokemonAbilitySapSipper        PokemonAbility = "SAP_SIPPER"
	PokemonAbilitySchooling        PokemonAbility = "SCHOOLING"
	PokemonAbilityScrappy          PokemonAbility = "SCRAPPY"
	PokemonAbilityScreenCleaner    PokemonAbility = "SCREEN_CLEANER"
	PokemonAbilitySeedSower        PokemonAbility = "SEED_SOWER"
	PokemonAbilitySereneGrace      PokemonAbility = "SERENE_GRACE"
	PokemonAbilityShadowShield     PokemonAbility = "SHADOW_SHIELD"
	PokemonAbilityShadowTag        PokemonAbility = "SHADOW_TAG"
	PokemonAbilitySharpness        PokemonAbility = "SHARPNESS"
	PokemonAbilityShedSkin         PokemonAbility = "SHED_SKIN"
	PokemonAbilitySheerForce       PokemonAbility = "SHEER_FORCE"
	PokemonAbilityShellArmor       PokemonAbility = "SHELL_ARMOR"
	PokemonAbilityShieldDust       PokemonAbility = "SHIELD_DUST"
	PokemonAbilityShieldsDown      PokemonAbility = "SHIELDS_DOWN"
	PokemonAbilitySimple           PokemonAbility = "SIMPLE"
	PokemonAbilitySkillLink        PokemonAbility = "SKILL_LINK"
	PokemonAbilitySlowStart        PokemonAbility = "SLOW_START"
	PokemonAbilitySlushRush        PokemonAbility = "SLUSH_RUSH"
	PokemonAbilitySniper           PokemonAbility = "SNIPER"
	PokemonAbilitySnowCloak        PokemonAbility = "SNOW_CLOAK"
	PokemonAbilitySnowWarning      PokemonAbility = "SNOW_WARNING"
	PokemonAbilitySolarPower       PokemonAbility = "SOLAR_POWER"
	PokemonAbilitySolidRock        PokemonAbility = "SOLID_ROCK"
	PokemonAbilitySoulHeart        PokemonAbility = "SOUL_HEART"
	PokemonAbilitySoundproof       PokemonAbility = "SOUNDPROOF"
	PokemonAbilitySpeedBoost       PokemonAbility = "SPEED_BOOST"
	PokemonAbilityStakeout         PokemonAbility = "STAKEOUT"
	PokemonAbilityStall            PokemonAbility = "STALL"
	PokemonAbilityStamina          PokemonAbility = "STAMINA"
	PokemonAbilityStanceChange     PokemonAbility = "STANCE_CHANGE"
	PokemonAbilityStatic           PokemonAbility = "STATIC"
	PokemonAbilitySteadfast        PokemonAbility = "STEADFAST"
	PokemonAbilitySteamEngine      PokemonAbility = "STEAM_ENGINE"
	PokemonAbilitySteelworker      PokemonAbility = "STEELWORKER"
	PokemonAbilityStench           PokemonAbility = "STENCH"
	PokemonAbilityStickyHold       PokemonAbility = "STICKY_HOLD"
	PokemonAbilityStormDrain       PokemonAbility = "STORM_DRAIN"
	PokemonAbilityStrongJaw        PokemonAbility = "STRONG_JAW"
	PokemonAbilitySturdy           PokemonAbility = "STURDY"
	PokemonAbilitySuctionCups      PokemonAbility = "SUCTION_CUPS"
	PokemonAbilitySuperLuck        PokemonAbility = "SUPER_LUCK"
	PokemonAbilitySupersweetSyrup  PokemonAbility = "SUPERSWEET_SYRUP"
	PokemonAbilitySupremeOverlord  PokemonAbility = "SUPREME_OVERLORD"
	PokemonAbilitySwarm            PokemonAbility = "SWARM"
	PokemonAbilitySweetVeil        PokemonAbility = "SWEET_VEIL"
	PokemonAbilitySwiftSwim        PokemonAbility = "SWIFT_SWIM"
	PokemonAbilitySwordOfRuin      PokemonAbility = "SWORD_OF_RUIN"
	PokemonAbilitySynchronize      PokemonAbility = "SYNCHRONIZE"
	PokemonAbilityTabletsOfRuin    PokemonAbility = "TABLETS_OF_RUIN"
	PokemonAbilityTangledFeet      PokemonAbility = "TANGLED_FEET"
	PokemonAbilityTechnician       PokemonAbility = "TECHNICIAN"
	PokemonAbilityTelepathy        PokemonAbility = "TELEPATHY"
	PokemonAbilityTeraShift        PokemonAbility = "TERA_SHIFT"
	PokemonAbilityTeravolt         PokemonAbility = "TERAVOLT"
	PokemonAbilityThermalExchange  PokemonAbility = "THERMAL_EXCHANGE"
	PokemonAbilityThickFat         PokemonAbility = "THICK_FAT"
	PokemonAbilityTintedLens       PokemonAbility = "TINTED_LENS"
	PokemonAbilityTorrent          PokemonAbility = "TORRENT"
	PokemonAbilityToughClaws       PokemonAbility = "TOUGH_CLAWS"
	PokemonAbilityToxicChain       PokemonAbility = "TOXIC_CHAIN"
	PokemonAbilityToxicDebris      PokemonAbility = "TOXIC_DEBRIS"
	PokemonAbilityTrace            PokemonAbility = "TRACE"
	PokemonAbilityTransistor       PokemonAbility = "TRANSISTOR"
	PokemonAbilityTriage           PokemonAbility = "TRIAGE"
	PokemonAbilityTruant           PokemonAbility = "TRUANT"
	PokemonAbilityTurboblaze       PokemonAbility = "TURBOBLAZE"
	PokemonAbilityUnaware          PokemonAbility = "UNAWARE"
	PokemonAbilityUnburden         PokemonAbility = "UNBURDEN"
	PokemonAbilityUnnerve          PokemonAbility = "UNNERVE"
	PokemonAbilityUnseenFist       PokemonAbility = "UNSEEN_FIST"
	PokemonAbilityVesselOfRuin     PokemonAbility = "VESSEL_OF_RUIN"
	PokemonAbilityVictoryStar      PokemonAbility = "VICTORY_STAR"
	PokemonAbilityVitalSpirit      PokemonAbility = "VITAL_SPIRIT"
	PokemonAbilityVoltAbsorb       PokemonAbility = "VOLT_ABSORB"
	PokemonAbilityWanderingSpirit  PokemonAbility = "WANDERING_SPIRIT"
	PokemonAbilityWaterAbsorb      PokemonAbility = "WATER_ABSORB"
	PokemonAbilityWaterBubble      PokemonAbility = "WATER_BUBBLE"
	PokemonAbilityWaterCompaction  PokemonAbility = "WATER_COMPACTION"
	PokemonAbilityWaterVeil        PokemonAbility = "WATER_VEIL"
	PokemonAbilityWeakArmor        PokemonAbility = "WEAK_ARMOR"
	PokemonAbilityWellBakedBody    PokemonAbility = "WELL_BAKED_BODY"
	PokemonAbilityWhiteSmoke       PokemonAbility = "WHITE_SMOKE"
	PokemonAbilityWimpOut          PokemonAbility = "WIMP_OUT"
	PokemonAbilityWindPower        PokemonAbility = "WIND_POWER"
	PokemonAbilityWindRider        PokemonAbility = "WIND_RIDER"
	PokemonAbilityWonderGuard      PokemonAbility = "WONDER_GUARD"
	PokemonAbilityWonderSkin       PokemonAbility = "WONDER_SKIN"
	PokemonAbilityZeroToHero       PokemonAbility = "ZERO_TO_HERO"
)

var AllPokemonAbility = []PokemonAbility{
	PokemonAbilityAdaptability,
	PokemonAbilityAftermath,
	PokemonAbilityAirLock,
	PokemonAbilityAngerPoint,
	PokemonAbilityAngerShell,
	PokemonAbilityAnticipation,
	PokemonAbilityArenaTrap,
	PokemonAbilityArmorTail,
	PokemonAbilityAromaVeil,
	PokemonAbilityAuraBreak,
	PokemonAbilityBadDreams,
	PokemonAbilityBallFetch,
	PokemonAbilityBattery,
	PokemonAbilityBattleArmor,
	PokemonAbilityBeadsOfRuin,
	PokemonAbilityBeastBoost,
	PokemonAbilityBerserk,
	PokemonAbilityBigPecks,
	PokemonAbilityBlaze,
	PokemonAbilityBulletproof,
	PokemonAbilityCheekPouch,
	PokemonAbilityChillingNeigh,
	PokemonAbilityChlorophyll,
	PokemonAbilityClearBody,
	PokemonAbilityCloudNine,
	PokemonAbilityColorChange,
	PokemonAbilityComatose,
	PokemonAbilityCommander,
	PokemonAbilityCompetitive,
	PokemonAbilityCompoundEyes,
	PokemonAbilityContrary,
	PokemonAbilityCorrosion,
	PokemonAbilityCottonDown,
	PokemonAbilityCudChew,
	PokemonAbilityCursedBody,
	PokemonAbilityCuteCharm,
	PokemonAbilityDamp,
	PokemonAbilityDancer,
	PokemonAbilityDark,
	PokemonAbilityAura,
	PokemonAbilityDauntlessShield,
	PokemonAbilityDazzling,
	PokemonAbilityDefeatist,
	PokemonAbilityDefiant,
	PokemonAbilityDisguise,
	PokemonAbilityDownload,
	PokemonAbilityDragonsMaw,
	PokemonAbilityDrizzle,
	PokemonAbilityDrought,
	PokemonAbilityDrySkin,
	PokemonAbilityEarlyBird,
	PokemonAbilityEarthEater,
	PokemonAbilityEffectSpore,
	PokemonAbilityElectricSurge,
	PokemonAbilityElectromorphosis,
	PokemonAbilityEmergencyExit,
	PokemonAbilityFairyAura,
	PokemonAbilityFilter,
	PokemonAbilityFlameBody,
	PokemonAbilityFlashFire,
	PokemonAbilityFlowerGift,
	PokemonAbilityFlowerVeil,
	PokemonAbilityFluffy,
	PokemonAbilityForecast,
	PokemonAbilityForewarn,
	PokemonAbilityFriendGuard,
	PokemonAbilityFrisk,
	PokemonAbilityFullMetalBody,
	PokemonAbilityFurCoat,
	PokemonAbilityGluttony,
	PokemonAbilityGoodAsGold,
	PokemonAbilityGooey,
	PokemonAbilityGrassySurge,
	PokemonAbilityGrimNeigh,
	PokemonAbilityGuardDog,
	PokemonAbilityGulpMissile,
	PokemonAbilityGuts,
	PokemonAbilityHadronEngine,
	PokemonAbilityHealer,
	PokemonAbilityHeatproof,
	PokemonAbilityHeavyMetal,
	PokemonAbilityHoneyGather,
	PokemonAbilityHospitality,
	PokemonAbilityHugePower,
	PokemonAbilityHungerSwitch,
	PokemonAbilityHustle,
	PokemonAbilityHydration,
	PokemonAbilityHyperCutter,
	PokemonAbilityIceBody,
	PokemonAbilityIceFace,
	PokemonAbilityIlluminate,
	PokemonAbilityIllusion,
	PokemonAbilityImmunity,
	PokemonAbilityInfiltrator,
	PokemonAbilityInnardsOut,
	PokemonAbilityInnerFocus,
	PokemonAbilityInsomnia,
	PokemonAbilityIntimidate,
	PokemonAbilityIntrepidSword,
	PokemonAbilityIronBarbs,
	PokemonAbilityIronFist,
	PokemonAbilityJustified,
	PokemonAbilityKeenEye,
	PokemonAbilityKlutz,
	PokemonAbilityLeafGuard,
	PokemonAbilityLevitate,
	PokemonAbilityLightMetal,
	PokemonAbilityLightningRod,
	PokemonAbilityLimber,
	PokemonAbilityLingeringAroma,
	PokemonAbilityLiquidOoze,
	PokemonAbilityMagicGuard,
	PokemonAbilityMagician,
	PokemonAbilityMagmaArmor,
	PokemonAbilityMagnetPull,
	PokemonAbilityMarvelScale,
	PokemonAbilityMegaLauncher,
	PokemonAbilityMerciless,
	PokemonAbilityMinus,
	PokemonAbilityMistySurge,
	PokemonAbilityMoldBreaker,
	PokemonAbilityMotorDrive,
	PokemonAbilityMoxie,
	PokemonAbilityMultitype,
	PokemonAbilityMummy,
	PokemonAbilityMyceliumMight,
	PokemonAbilityNaturalCure,
	PokemonAbilityNeutralizingGas,
	PokemonAbilityNoGuard,
	PokemonAbilityNormalize,
	PokemonAbilityOblivious,
	PokemonAbilityOpportunist,
	PokemonAbilityOrichalcumPulse,
	PokemonAbilityOvercoat,
	PokemonAbilityOvergrow,
	PokemonAbilityOwnTempo,
	PokemonAbilityPickpocket,
	PokemonAbilityPickup,
	PokemonAbilityPlus,
	PokemonAbilityPoisonHeal,
	PokemonAbilityPoisonPoint,
	PokemonAbilityPoisonPuppeteer,
	PokemonAbilityPoisonTouch,
	PokemonAbilityPowerSpot,
	PokemonAbilityPrankster,
	PokemonAbilityPressure,
	PokemonAbilityPrismArmor,
	PokemonAbilityProtosynthesis,
	PokemonAbilityPsychicSurge,
	PokemonAbilityPunkRock,
	PokemonAbilityPurePower,
	PokemonAbilityPurifyingSalt,
	PokemonAbilityQuarkDrive,
	PokemonAbilityQueenlyMajesty,
	PokemonAbilityQuickFeet,
	PokemonAbilityRksSystem,
	PokemonAbilityRainDish,
	PokemonAbilityRattled,
	PokemonAbilityReceiver,
	PokemonAbilityReckless,
	PokemonAbilityRefrigerate,
	PokemonAbilityRegenerator,
	PokemonAbilityRipen,
	PokemonAbilityRivalry,
	PokemonAbilityRockHead,
	PokemonAbilityRoughSkin,
	PokemonAbilityRunAway,
	PokemonAbilitySandForce,
	PokemonAbilitySandRush,
	PokemonAbilitySandSpit,
	PokemonAbilitySandStream,
	PokemonAbilitySandVeil,
	PokemonAbilitySapSipper,
	PokemonAbilitySchooling,
	PokemonAbilityScrappy,
	PokemonAbilityScreenCleaner,
	PokemonAbilitySeedSower,
	PokemonAbilitySereneGrace,
	PokemonAbilityShadowShield,
	PokemonAbilityShadowTag,
	PokemonAbilitySharpness,
	PokemonAbilityShedSkin,
	PokemonAbilitySheerForce,
	PokemonAbilityShellArmor,
	PokemonAbilityShieldDust,
	PokemonAbilityShieldsDown,
	PokemonAbilitySimple,
	PokemonAbilitySkillLink,
	PokemonAbilitySlowStart,
	PokemonAbilitySlushRush,
	PokemonAbilitySniper,
	PokemonAbilitySnowCloak,
	PokemonAbilitySnowWarning,
	PokemonAbilitySolarPower,
	PokemonAbilitySolidRock,
	PokemonAbilitySoulHeart,
	PokemonAbilitySoundproof,
	PokemonAbilitySpeedBoost,
	PokemonAbilityStakeout,
	PokemonAbilityStall,
	PokemonAbilityStamina,
	PokemonAbilityStanceChange,
	PokemonAbilityStatic,
	PokemonAbilitySteadfast,
	PokemonAbilitySteamEngine,
	PokemonAbilitySteelworker,
	PokemonAbilityStench,
	PokemonAbilityStickyHold,
	PokemonAbilityStormDrain,
	PokemonAbilityStrongJaw,
	PokemonAbilitySturdy,
	PokemonAbilitySuctionCups,
	PokemonAbilitySuperLuck,
	PokemonAbilitySupersweetSyrup,
	PokemonAbilitySupremeOverlord,
	PokemonAbilitySwarm,
	PokemonAbilitySweetVeil,
	PokemonAbilitySwiftSwim,
	PokemonAbilitySwordOfRuin,
	PokemonAbilitySynchronize,
	PokemonAbilityTabletsOfRuin,
	PokemonAbilityTangledFeet,
	PokemonAbilityTechnician,
	PokemonAbilityTelepathy,
	PokemonAbilityTeraShift,
	PokemonAbilityTeravolt,
	PokemonAbilityThermalExchange,
	PokemonAbilityThickFat,
	PokemonAbilityTintedLens,
	PokemonAbilityTorrent,
	PokemonAbilityToughClaws,
	PokemonAbilityToxicChain,
	PokemonAbilityToxicDebris,
	PokemonAbilityTrace,
	PokemonAbilityTransistor,
	PokemonAbilityTriage,
	PokemonAbilityTruant,
	PokemonAbilityTurboblaze,
	PokemonAbilityUnaware,
	PokemonAbilityUnburden,
	PokemonAbilityUnnerve,
	PokemonAbilityUnseenFist,
	PokemonAbilityVesselOfRuin,
	PokemonAbilityVictoryStar,
	PokemonAbilityVitalSpirit,
	PokemonAbilityVoltAbsorb,
	PokemonAbilityWanderingSpirit,
	PokemonAbilityWaterAbsorb,
	PokemonAbilityWaterBubble,
	PokemonAbilityWaterCompaction,
	PokemonAbilityWaterVeil,
	PokemonAbilityWeakArmor,
	PokemonAbilityWellBakedBody,
	PokemonAbilityWhiteSmoke,
	PokemonAbilityWimpOut,
	PokemonAbilityWindPower,
	PokemonAbilityWindRider,
	PokemonAbilityWonderGuard,
	PokemonAbilityWonderSkin,
	PokemonAbilityZeroToHero,
}

func (e PokemonAbility) IsValid() bool {
	switch e {
	case PokemonAbilityAdaptability, PokemonAbilityAftermath, PokemonAbilityAirLock, PokemonAbilityAngerPoint, PokemonAbilityAngerShell, PokemonAbilityAnticipation, PokemonAbilityArenaTrap, PokemonAbilityArmorTail, PokemonAbilityAromaVeil, PokemonAbilityAuraBreak, PokemonAbilityBadDreams, PokemonAbilityBallFetch, PokemonAbilityBattery, PokemonAbilityBattleArmor, PokemonAbilityBeadsOfRuin, PokemonAbilityBeastBoost, PokemonAbilityBerserk, PokemonAbilityBigPecks, PokemonAbilityBlaze, PokemonAbilityBulletproof, PokemonAbilityCheekPouch, PokemonAbilityChillingNeigh, PokemonAbilityChlorophyll, PokemonAbilityClearBody, PokemonAbilityCloudNine, PokemonAbilityColorChange, PokemonAbilityComatose, PokemonAbilityCommander, PokemonAbilityCompetitive, PokemonAbilityCompoundEyes, PokemonAbilityContrary, PokemonAbilityCorrosion, PokemonAbilityCottonDown, PokemonAbilityCudChew, PokemonAbilityCursedBody, PokemonAbilityCuteCharm, PokemonAbilityDamp, PokemonAbilityDancer, PokemonAbilityDark, PokemonAbilityAura, PokemonAbilityDauntlessShield, PokemonAbilityDazzling, PokemonAbilityDefeatist, PokemonAbilityDefiant, PokemonAbilityDisguise, PokemonAbilityDownload, PokemonAbilityDragonsMaw, PokemonAbilityDrizzle, PokemonAbilityDrought, PokemonAbilityDrySkin, PokemonAbilityEarlyBird, PokemonAbilityEarthEater, PokemonAbilityEffectSpore, PokemonAbilityElectricSurge, PokemonAbilityElectromorphosis, PokemonAbilityEmergencyExit, PokemonAbilityFairyAura, PokemonAbilityFilter, PokemonAbilityFlameBody, PokemonAbilityFlashFire, PokemonAbilityFlowerGift, PokemonAbilityFlowerVeil, PokemonAbilityFluffy, PokemonAbilityForecast, PokemonAbilityForewarn, PokemonAbilityFriendGuard, PokemonAbilityFrisk, PokemonAbilityFullMetalBody, PokemonAbilityFurCoat, PokemonAbilityGluttony, PokemonAbilityGoodAsGold, PokemonAbilityGooey, PokemonAbilityGrassySurge, PokemonAbilityGrimNeigh, PokemonAbilityGuardDog, PokemonAbilityGulpMissile, PokemonAbilityGuts, PokemonAbilityHadronEngine, PokemonAbilityHealer, PokemonAbilityHeatproof, PokemonAbilityHeavyMetal, PokemonAbilityHoneyGather, PokemonAbilityHospitality, PokemonAbilityHugePower, PokemonAbilityHungerSwitch, PokemonAbilityHustle, PokemonAbilityHydration, PokemonAbilityHyperCutter, PokemonAbilityIceBody, PokemonAbilityIceFace, PokemonAbilityIlluminate, PokemonAbilityIllusion, PokemonAbilityImmunity, PokemonAbilityInfiltrator, PokemonAbilityInnardsOut, PokemonAbilityInnerFocus, PokemonAbilityInsomnia, PokemonAbilityIntimidate, PokemonAbilityIntrepidSword, PokemonAbilityIronBarbs, PokemonAbilityIronFist, PokemonAbilityJustified, PokemonAbilityKeenEye, PokemonAbilityKlutz, PokemonAbilityLeafGuard, PokemonAbilityLevitate, PokemonAbilityLightMetal, PokemonAbilityLightningRod, PokemonAbilityLimber, PokemonAbilityLingeringAroma, PokemonAbilityLiquidOoze, PokemonAbilityMagicGuard, PokemonAbilityMagician, PokemonAbilityMagmaArmor, PokemonAbilityMagnetPull, PokemonAbilityMarvelScale, PokemonAbilityMegaLauncher, PokemonAbilityMerciless, PokemonAbilityMinus, PokemonAbilityMistySurge, PokemonAbilityMoldBreaker, PokemonAbilityMotorDrive, PokemonAbilityMoxie, PokemonAbilityMultitype, PokemonAbilityMummy, PokemonAbilityMyceliumMight, PokemonAbilityNaturalCure, PokemonAbilityNeutralizingGas, PokemonAbilityNoGuard, PokemonAbilityNormalize, PokemonAbilityOblivious, PokemonAbilityOpportunist, PokemonAbilityOrichalcumPulse, PokemonAbilityOvercoat, PokemonAbilityOvergrow, PokemonAbilityOwnTempo, PokemonAbilityPickpocket, PokemonAbilityPickup, PokemonAbilityPlus, PokemonAbilityPoisonHeal, PokemonAbilityPoisonPoint, PokemonAbilityPoisonPuppeteer, PokemonAbilityPoisonTouch, PokemonAbilityPowerSpot, PokemonAbilityPrankster, PokemonAbilityPressure, PokemonAbilityPrismArmor, PokemonAbilityProtosynthesis, PokemonAbilityPsychicSurge, PokemonAbilityPunkRock, PokemonAbilityPurePower, PokemonAbilityPurifyingSalt, PokemonAbilityQuarkDrive, PokemonAbilityQueenlyMajesty, PokemonAbilityQuickFeet, PokemonAbilityRksSystem, PokemonAbilityRainDish, PokemonAbilityRattled, PokemonAbilityReceiver, PokemonAbilityReckless, PokemonAbilityRefrigerate, PokemonAbilityRegenerator, PokemonAbilityRipen, PokemonAbilityRivalry, PokemonAbilityRockHead, PokemonAbilityRoughSkin, PokemonAbilityRunAway, PokemonAbilitySandForce, PokemonAbilitySandRush, PokemonAbilitySandSpit, PokemonAbilitySandStream, PokemonAbilitySandVeil, PokemonAbilitySapSipper, PokemonAbilitySchooling, PokemonAbilityScrappy, PokemonAbilityScreenCleaner, PokemonAbilitySeedSower, PokemonAbilitySereneGrace, PokemonAbilityShadowShield, PokemonAbilityShadowTag, PokemonAbilitySharpness, PokemonAbilityShedSkin, PokemonAbilitySheerForce, PokemonAbilityShellArmor, PokemonAbilityShieldDust, PokemonAbilityShieldsDown, PokemonAbilitySimple, PokemonAbilitySkillLink, PokemonAbilitySlowStart, PokemonAbilitySlushRush, PokemonAbilitySniper, PokemonAbilitySnowCloak, PokemonAbilitySnowWarning, PokemonAbilitySolarPower, PokemonAbilitySolidRock, PokemonAbilitySoulHeart, PokemonAbilitySoundproof, PokemonAbilitySpeedBoost, PokemonAbilityStakeout, PokemonAbilityStall, PokemonAbilityStamina, PokemonAbilityStanceChange, PokemonAbilityStatic, PokemonAbilitySteadfast, PokemonAbilitySteamEngine, PokemonAbilitySteelworker, PokemonAbilityStench, PokemonAbilityStickyHold, PokemonAbilityStormDrain, PokemonAbilityStrongJaw, PokemonAbilitySturdy, PokemonAbilitySuctionCups, PokemonAbilitySuperLuck, PokemonAbilitySupersweetSyrup, PokemonAbilitySupremeOverlord, PokemonAbilitySwarm, PokemonAbilitySweetVeil, PokemonAbilitySwiftSwim, PokemonAbilitySwordOfRuin, PokemonAbilitySynchronize, PokemonAbilityTabletsOfRuin, PokemonAbilityTangledFeet, PokemonAbilityTechnician, PokemonAbilityTelepathy, PokemonAbilityTeraShift, PokemonAbilityTeravolt, PokemonAbilityThermalExchange, PokemonAbilityThickFat, PokemonAbilityTintedLens, PokemonAbilityTorrent, PokemonAbilityToughClaws, PokemonAbilityToxicChain, PokemonAbilityToxicDebris, PokemonAbilityTrace, PokemonAbilityTransistor, PokemonAbilityTriage, PokemonAbilityTruant, PokemonAbilityTurboblaze, PokemonAbilityUnaware, PokemonAbilityUnburden, PokemonAbilityUnnerve, PokemonAbilityUnseenFist, PokemonAbilityVesselOfRuin, PokemonAbilityVictoryStar, PokemonAbilityVitalSpirit, PokemonAbilityVoltAbsorb, PokemonAbilityWanderingSpirit, PokemonAbilityWaterAbsorb, PokemonAbilityWaterBubble, PokemonAbilityWaterCompaction, PokemonAbilityWaterVeil, PokemonAbilityWeakArmor, PokemonAbilityWellBakedBody, PokemonAbilityWhiteSmoke, PokemonAbilityWimpOut, PokemonAbilityWindPower, PokemonAbilityWindRider, PokemonAbilityWonderGuard, PokemonAbilityWonderSkin, PokemonAbilityZeroToHero:
		return true
	}
	return false
}

func (e PokemonAbility) String() string {
	return string(e)
}

func (e *PokemonAbility) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PokemonAbility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PokemonAbility", str)
	}
	return nil
}

func (e PokemonAbility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *PokemonAbility) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e PokemonAbility) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}

type PokemonType string

const (
	PokemonTypeBug      PokemonType = "BUG"
	PokemonTypeDragon   PokemonType = "DRAGON"
	PokemonTypeFairy    PokemonType = "FAIRY"
	PokemonTypeFire     PokemonType = "FIRE"
	PokemonTypeGhost    PokemonType = "GHOST"
	PokemonTypeGround   PokemonType = "GROUND"
	PokemonTypeNormal   PokemonType = "NORMAL"
	PokemonTypePsychic  PokemonType = "PSYCHIC"
	PokemonTypeSteel    PokemonType = "STEEL"
	PokemonTypeDark     PokemonType = "DARK"
	PokemonTypeElectric PokemonType = "ELECTRIC"
	PokemonTypeFighting PokemonType = "FIGHTING"
	PokemonTypeFlying   PokemonType = "FLYING"
	PokemonTypeGrass    PokemonType = "GRASS"
	PokemonTypeIce      PokemonType = "ICE"
	PokemonTypePoison   PokemonType = "POISON"
	PokemonTypeRock     PokemonType = "ROCK"
	PokemonTypeWater    PokemonType = "WATER"
)

var AllPokemonType = []PokemonType{
	PokemonTypeBug,
	PokemonTypeDragon,
	PokemonTypeFairy,
	PokemonTypeFire,
	PokemonTypeGhost,
	PokemonTypeGround,
	PokemonTypeNormal,
	PokemonTypePsychic,
	PokemonTypeSteel,
	PokemonTypeDark,
	PokemonTypeElectric,
	PokemonTypeFighting,
	PokemonTypeFlying,
	PokemonTypeGrass,
	PokemonTypeIce,
	PokemonTypePoison,
	PokemonTypeRock,
	PokemonTypeWater,
}

func (e PokemonType) IsValid() bool {
	switch e {
	case PokemonTypeBug, PokemonTypeDragon, PokemonTypeFairy, PokemonTypeFire, PokemonTypeGhost, PokemonTypeGround, PokemonTypeNormal, PokemonTypePsychic, PokemonTypeSteel, PokemonTypeDark, PokemonTypeElectric, PokemonTypeFighting, PokemonTypeFlying, PokemonTypeGrass, PokemonTypeIce, PokemonTypePoison, PokemonTypeRock, PokemonTypeWater:
		return true
	}
	return false
}

func (e PokemonType) String() string {
	return string(e)
}

func (e *PokemonType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PokemonType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PokemonType", str)
	}
	return nil
}

func (e PokemonType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *PokemonType) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalGQL(s)
}

func (e PokemonType) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	e.MarshalGQL(&buf)
	return buf.Bytes(), nil
}
