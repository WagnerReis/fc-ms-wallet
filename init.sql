-- Script de inicialização do banco de dados wallet
-- Cria as tabelas e alguns dados de exemplo

-- Criação das tabelas
CREATE TABLE IF NOT EXISTS clients (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL,
    balance DECIMAL(10,2) DEFAULT 0.00,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    account_id_from VARCHAR(255) NOT NULL,
    account_id_to VARCHAR(255) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id_from) REFERENCES accounts(id),
    FOREIGN KEY (account_id_to) REFERENCES accounts(id)
);

-- Inserção de dados de exemplo
INSERT INTO clients (id, name, email, created_at) VALUES 
('client-001', 'João Silva', 'joao@email.com', NOW()),
('client-002', 'Maria Santos', 'maria@email.com', NOW()),
('client-003', 'Pedro Oliveira', 'pedro@email.com', NOW());

INSERT INTO accounts (id, client_id, balance, created_at) VALUES 
('account-001', 'client-001', 1000.00, NOW()),
('account-002', 'client-002', 500.00, NOW()),
('account-003', 'client-003', 750.00, NOW());

-- Inserção de uma transação de exemplo
INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES 
('trans-001', 'account-001', 'account-002', 100.00, NOW());
