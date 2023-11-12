use snowdon::{Epoch, MachineId};

pub struct SnowflakeParameters;

impl Epoch for SnowflakeParameters {
    fn millis_since_unix() -> u64 {
        1420070400000
    }
}

impl MachineId for SnowflakeParameters {
    fn machine_id() -> u64 {
        0
    }
}

pub struct SnowflakeProvider {
    pub generator:
        snowdon::Generator<snowdon::ClassicLayout<SnowflakeParameters>, SnowflakeParameters>,
}

impl SnowflakeProvider {
    pub fn generate(&self) -> String {
        self.generator.generate().unwrap().to_string()
    }
}
