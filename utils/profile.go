func Profile(fileName string) {
	f, err := os.Create(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}