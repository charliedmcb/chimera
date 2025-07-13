package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var (
	prefixes []string = []string{
		"Tag Me",
		"Core Me",
		"Rush",
		"POG",
		"The All Seeing",
		"Flatlined",
		"Kill",
		"FA",
		"Glacier",
		"DJ",
		"Banned",
		"OP",
		"Bad Pub",
		"Disgraced",
		"Accelerated",
	}

	corpIdPrefixes []string = []string{
		"Acme Consulting",
		"AgInfusion",
		"Ampère",
		"Argus Security",
		"Asa Group",
		"Azmari EdTech",
		"Blue Sun",
		"Cerebral Imaging",
		"Chronos Protocol",
		"Custom Biotics",
		"Cyber Bureau",
		"Cybernetics Division",
		"Earth Station",
		"Fringe Applications",
		"Gagarin Deep Space",
		"GameNET",
		"GRNDL",
		"Haarpsichord Studios",
		"Haas-Bioroid",
		"Harishchandra Ent.",
		"Harmony Medtech",
		"Hyoubu Institute",
		"Industrial Genomics",
		"Information Dynamics",
		"Issuaq Adaptics",
		"Jemison Astronautics",
		"Jinteki Biotech",
		"Jinteki",
		"MirrorMorph",
		"Mti Mwekundu",
		"NBN",
		"Near-Earth Hub",
		"New Angeles Sol",
		"NEXT Design",
		"Nisei Division",
		"Ob Superheavy Logistics",
		"Pālanā Foods",
		"Pravdivost Consulting",
		"Saraswati Mnemonics",
		"Seidr Laboratories",
		"Skorpios Defense Systems",
		"Spark Agency",
		"Sportsmetal",
		"SSO Industries",
		"Strategic Innovations",
		"SYNC",
		"Synthetic Systems",
		"Tennin Institute",
		"The Foundry",
		"The Outfit",
		"The Shadow",
		"The Syndicate",
		"Thule Subsea",
		"Titan Transnational",
		"Weyland Consortium",
		"A Teia",
		"Epiphany Analytica",
	}

	corpIdSuffixes []string = []string{
		"The Truth You Need",
		"New Miracles for a New World",
		"Cybernetics For Anyone",
		"Protection Guaranteed",
		"Security Through Vigilance",
		"Shaping the Future",
		"Powering the Future",
		"Infinite Frontiers",
		"Selective Mind-mapping",
		"Engineered for Success",
		"Keeping the Peace",
		"Humanity Upgraded",
		"SEA Headquarters",
		"Tomorrow, Today",
		"Expanding the Horizon",
		"Where Dreams are Real",
		"Power Unleashed",
		"Entertainment Unleashed",
		"Architects of Tomorrow",
		"Engineering the Future",
		"Precision Design",
		"Stronger Together",
		"Where You're the Star",
		"Biomedical Pioneer",
		"Absolute Clarity",
		"Growing Solutions",
		"All You Need To Know",
		"Sustaining Diversity",
		"Sacrifice. Audacity. Success.",
		"Life Imagined",
		"Personal Evolution",
		"Potential Unleashed",
		"Replicating Perfection",
		"Restoring Humanity",
		"Endless Iteration",
		"Life Improved",
		"Controlling the Message",
		"Making News",
		"Reality Plus",
		"The World is Yours*",
		"Broadcast Center",
		"Your News",
		"Guarding the Net",
		"The Next Generation",
		"Extract. Export. Excel.",
		"Sustainable Growth",
		"Political Solutions",
		"Endless Exploration",
		"Destiny Defined",
		"Persuasive Power",
		"Worldswide Reach",
		"Go Big or Go Home",
		"Fueling Innovation",
		"Future Forward",
		"Everything, Everywhere",
		"The World Re-imagined",
		"The Secrets Within",
		"Refining the Process",
		"Family Owned and Operated",
		"Pulling the Strings",
		"Profit over Principle",
		"Safety Below",
		"Investing In Your Future",
		"Because We Built It",
		"Builder of Nations",
		"Building a Better World",
		"Built to Last",
		"IP Recovery",
		"Nations Undivided",
	}

	runnerIdPrefixes []string = []string{
		"419",
		"Adam",
		"Akiko Nisei",
		"Alice Merchant",
		"Andromeda",
		"Apex",
		"Armand \"Geist\" Walker",
		"Ayla \"Bios\" Rahim",
		"Az McCaffrey",
		"Boris \"Syfr\" Kovac",
		"Captain Padma Isbister",
		"Chaos Theory",
		"Edward Kim",
		"Ele \"Smoke\" Scovak",
		"Esâ Afontov",
		"Exile",
		"Freedom Khumalo",
		"Gabriel Santiago",
		"Hayley Kaplan",
		"Hoshiko Shiro",
		"Iain Stirling",
		"Jamie \"Bzzz\" Micken",
		"Jesminder Sareen",
		"Kabonesa Wu",
		"Kate \"Mac\" McCaffrey",
		"Ken \"Express\" Tenma",
		"Khan",
		"Laramy Fisk",
		"Lat",
		"Leela Patel",
		"Liza Talking Thunder",
		"Los",
		"MaxX",
		"Nasir Meidan",
		"Nathaniel \"Gnat\" Hall",
		"Nero Severn",
		"Noise",
		"Nova Initiumia",
		"Null",
		"Nyusha \"Sable\" Sintashta",
		"Omar Keung",
		"Quetzal",
		"Reina Roja",
		"René \"Loup\" Arcemont",
		"Rielle \"Kit\" Peddler",
		"Silhouette",
		"Steve Cambridge",
		"Sunny Lebeau",
		"Tāo Salonga",
		"The Catalyst",
		"The Masque",
		"The Professor",
		"Valencia Estevez",
		"Whizzard",
		"Wyvern",
		"Zahya Sadeghi",
		"Arissana Rocha Nahu",
		"Mercury",
	}

	runnerIdSuffixes []string = []string{
		"Amoral Scammer",
		"Compulsive Hacker",
		"Head Case",
		"Clan Agitator",
		"Dispossessed Ristie",
		"Invasive Predator",
		"Tech Lord",
		"Simulant Specialist",
		"Mechanical Prodigy",
		"Crafty Veteran",
		"Intrepid Explorer",
		"Wünderkind",
		"Humanity's Hammer",
		"Cynosure of the Net",
		"Eco-Insurrectionist",
		"Streethawk",
		"Crypto-Anarchist",
		"Consummate Professional",
		"Universal Scholar",
		"Untold Protagonist",
		"Retired Spook",
		"Techno Savant",
		"Girl Behind the Curtain",
		"Netspace Thrillseeker",
		"Digital Tinker",
		"Disappeared Clone",
		"Savvy Skiptracer",
		"Savvy Investor",
		"Ethical Freelancer",
		"Trained Pragmatist",
		"Prominent Legislator",
		"Data Hijacker",
		"Maximum Punk Rock",
		"Cyber Explorer",
		"One-of-a-Kind",
		"Information Broker",
		"Hacker Extraordinaire",
		"Catalyst & Impetus",
		"Whistleblower",
		"Symphonic Prodigy",
		"Conspiracy Theorist",
		"Free Spirit",
		"Freedom Fighter",
		"Party Animal",
		"Transhuman",
		"Stealth Operative",
		"Master Grifter",
		"Security Specialist",
		"Telepresence Magician",
		"Convention Breaker",
		"Cyber General",
		"Keeper of Knowledge",
		"The Angel of Cayambe",
		"Master Gamer",
		"Chemically Enhanced",
		"Versatile Smuggler",
		"Street Artist",
		"Chrome Libertador",
	}

	suffixes []string = []string{
		"Ice",
		"Bioroids",
		"Janus",
		"Jackson Howard",
		"Estelle Moon",
		"Pancakes",
		"Wyldcakes",
		"DJ Steve",
		"AstroScript",
		"Beth",
		"Wheels",
		"Bones",
		"Wu",
		"Ag",
		"Keeling",
	}
)

// generateNameFromSeed creates a deck name from a numeric seed
func generateNameFromSeed(seedValue int64) string {
	rand.Seed(seedValue)
	prefix := prefixes[rand.Intn(len(prefixes))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	name := fmt.Sprintf("%s %s (%d)", prefix, suffix, seedValue%1000)
	return name
}

// hashStringToSeed converts a string to a deterministic seed
func hashStringToSeed(input string) int64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(input))
	return int64(hasher.Sum64())
}

func generateDeckNameAndSeed() (string, int64) {
	seed := time.Now().UnixNano()
	name := generateNameFromSeed(seed)
	return name, hashStringToSeed(name)
}

func generateDeckNameAndSeedFromInput(input string) (string, int64) {
	// Check if input matches the format "prefix suffix (X)"
	pattern := regexp.MustCompile(`^(.+) (.+) \((\d+)\)$`)
	matches := pattern.FindStringSubmatch(input)
	
	if matches != nil && len(matches) == 4 {
		prefix := matches[1]
		suffix := matches[2]
		numberStr := matches[3]
		
		// Check if prefix and suffix are valid
		validPrefix := false
		for _, p := range prefixes {
			if p == prefix {
				validPrefix = true
				break
			}
		}
		
		validSuffix := false
		for _, s := range suffixes {
			if s == suffix {
				validSuffix = true
				break
			}
		}
		
		if validPrefix && validSuffix {
			number, err := strconv.Atoi(numberStr)
			if err == nil && number < 1000 {
				// Use the input name as-is and hash it for the seed
				return input, hashStringToSeed(input)
			}
		}
	}
	
	// If input doesn't match format or isn't valid, hash it and generate a new name
	seed := hashStringToSeed(input)
	name := generateNameFromSeed(seed)
	return name, hashStringToSeed(name)
}
