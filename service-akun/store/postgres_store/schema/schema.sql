CREATE TABLE "nasabah" (
  "nasabah_id" bigserial PRIMARY KEY,
  "nama" varchar NOT NULL,
  "nik" varchar UNIQUE NOT NULL,
  "no_hp" varchar UNIQUE NOT NULL,
  "tgl_dibuat" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "akun" (
  "akun_id" bigserial PRIMARY KEY,
  "nasabah_id" bigserial NOT NULL,
  "no_rekening" varchar UNIQUE NOT NULL,
  "saldo" bigint NOT NULL,
  "tgl_dibuat" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "mutasi" (
  "mutasi_id" bigserial PRIMARY KEY,
  "no_rekening" varchar NOT NULL,
  "kode_transaksi" varchar NOT NULL,
  "nominal" bigint NOT NULL,
  "waktu" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "akun" ADD FOREIGN KEY ("nasabah_id") REFERENCES "nasabah" ("nasabah_id");

ALTER TABLE "mutasi" ADD FOREIGN KEY ("no_rekening") REFERENCES "akun" ("no_rekening");
