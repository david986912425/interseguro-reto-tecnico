import {NextFunction, Request, Response} from "express";

import jwt from "jsonwebtoken";
import {env} from "../utils/env";

export default function validateJWT(req: Request, res: Response, next: NextFunction): void {
    try {
        const authHeader = req.header("Authorization");
        if (!authHeader) {
            res.status(401).json({
                ok: false,
                msg: "No hay token en la petición",
            });
            return;
        }

        const parts = authHeader.split(" ");
        if (parts.length !== 2 || parts[0] !== "Bearer") {
            res.status(401).json({
                ok: false,
                msg: "Formato de token inválido",
            });
            return;
        }

        const token = parts[1];

        const payload = jwt.verify(token, env.JWT_KEY);

        if (typeof payload !== "object" || !payload.email) {
            res.status(401).json({
                ok: false,
                msg: "Token inválido",
            });
            return;
        }

        (req as any).email = payload.email;

        next();
    } catch (error) {
        res.status(401).json({
            ok: false,
            msg: "Token no válido",
        });
    }
}