package v1

// type Profiler struct {
// 	config     *config.Config
// 	collector  *collector.ProfileCollector
// 	uploader   *uploader.ProfileUploader
// 	queue      chan *collector.ProfileData
// 	stopChan   chan struct{}
// 	wg         sync.WaitGroup
// 	logger     utils.Logger
// 	ctx        context.Context
// 	cancelFunc context.CancelFunc
// }

// func NewProfiler(cfg *config.Config) (*Profiler, error) {
// 	if err := cfg.Validate(); err != nil {
// 		return nil, err
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	return &Profiler{
// 		config:     cfg,
// 		collector:  collector.NewProfileCollector(),
// 		uploader:   uploader.NewUploader(cfg.Endpoint, cfg.Token),
// 		queue:      make(chan *collector.ProfileData, cfg.MaxQueueSize),
// 		stopChan:   make(chan struct{}),
// 		logger:     utils.NewDefaultLogger(cfg.Debug),
// 		ctx:        ctx,
// 		cancelFunc: cancel,
// 	}, nil
// }

// func (p *Profiler) Start() error {
// 	p.wg.Add(2)
// 	go p.collectionLoop()
// 	go p.uploadLoop()
// 	return nil
// }

// func (p *Profiler) Stop() {
// 	p.cancelFunc()
// 	close(p.stopChan)
// 	p.wg.Wait()
// }

// func (p *Profiler) collectionLoop() {
// 	defer p.wg.Done()

// 	ticker := time.NewTicker(p.config.Interval)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			p.collectAndQueue()
// 		case <-p.ctx.Done():
// 			return
// 		}
// 	}
// }

// func (p *Profiler) collectAndQueue() {
// 	for _, pt := range p.config.ProfileTypes {
// 		data, err := p.collector.Collect(collector.ProfileType(pt), p.config.Interval)
// 		if err != nil {
// 			p.logger.Error("Collection failed", "type", pt, "error", err)
// 			continue
// 		}

// 		select {
// 		case p.queue <- data:
// 		default:
// 			p.logger.Warn("Queue full, dropping profile", "type", pt)
// 		}
// 	}
// }

// func (p *Profiler) uploadLoop() {
// 	defer p.wg.Done()

// 	for {
// 		select {
// 		case data := <-p.queue:
// 			if err := p.uploadWithRetry(data); err != nil {
// 				p.logger.Error("Final upload failed", "type", data.Type, "error", err)
// 			}
// 		case <-p.ctx.Done():
// 			return
// 		}
// 	}
// }

// func (p *Profiler) uploadWithRetry(data *collector.ProfileData) error {
// 	const maxRetries = 3
// 	var err error

// 	for i := 0; i < maxRetries; i++ {
// 		if err = p.uploader.Upload(data); err == nil {
// 			return nil
// 		}
// 		time.Sleep(time.Duration(i+1) * time.Second)
// 	}
// 	return err
// }
