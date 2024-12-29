CREATE TABLE IF NOT EXISTS wallet (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  total integer DEFAULT 0
);
