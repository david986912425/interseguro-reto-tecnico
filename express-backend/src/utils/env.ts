import dotenv from 'dotenv';
dotenv.config();

function getEnvVar(name: string): string {
    const value = process.env[name];
    if (!value) {
        throw new Error(`La variable de entorno ${name} no está definida`);
    }
    return value;
}

export const env = {
    MONGO_URI: getEnvVar('MONGO_URI'),
    JWT_KEY: getEnvVar('JWT_KEY'),
};
