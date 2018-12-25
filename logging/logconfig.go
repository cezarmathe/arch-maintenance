package logging

type configLogger struct{}

func (c *configLogger) Trace() {}

func (c *configLogger) Debug() {}

func (c *configLogger) Info() {}

func (c *configLogger) Warn() {}

func (c *configLogger) Error() {}

func (c *configLogger) Fatal() {}

func (c *configLogger) Panic() {}
