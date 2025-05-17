import mongoose from 'mongoose';
import app from './app';
import { env } from './utils/env';

const PORT = process.env.PORT || 3000;

async function main() {
    try {
        await mongoose.connect(env.MONGO_URI);

        console.log('✅ Conectado a MongoDB');

        app.listen(PORT, () => {
            console.log(`🚀 Server is running on http://localhost:${PORT}`);
        });
    } catch (error) {
        console.error('❌ Error conectando a MongoDB:', error);
        process.exit(1);
    }
}

main();
