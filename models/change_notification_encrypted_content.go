package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeNotificationEncryptedContent 
type ChangeNotificationEncryptedContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Base64-encoded encrypted data that produces a full resource respresented as JSON. The data has been encrypted with the provided dataKey using an AES/CBC/PKCS5PADDING cipher suite.
    data *string
    // Base64-encoded symmetric key generated by Microsoft Graph to encrypt the data value and to generate the data signature. This key is encrypted with the certificate public key that was provided during the subscription. It must be decrypted with the certificate private key before it can be used to decrypt the data or verify the signature. This key has been encrypted with the following cipher suite: RSA/ECB/OAEPWithSHA1AndMGF1Padding.
    dataKey *string
    // Base64-encoded HMAC-SHA256 hash of the data for validation purposes.
    dataSignature *string
    // ID of the certificate used to encrypt the dataKey.
    encryptionCertificateId *string
    // Hexadecimal representation of the thumbprint of the certificate used to encrypt the dataKey.
    encryptionCertificateThumbprint *string
    // The OdataType property
    odataType *string
}
// NewChangeNotificationEncryptedContent instantiates a new changeNotificationEncryptedContent and sets the default values.
func NewChangeNotificationEncryptedContent()(*ChangeNotificationEncryptedContent) {
    m := &ChangeNotificationEncryptedContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChangeNotificationEncryptedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChangeNotificationEncryptedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeNotificationEncryptedContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeNotificationEncryptedContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. Base64-encoded encrypted data that produces a full resource respresented as JSON. The data has been encrypted with the provided dataKey using an AES/CBC/PKCS5PADDING cipher suite.
func (m *ChangeNotificationEncryptedContent) GetData()(*string) {
    return m.data
}
// GetDataKey gets the dataKey property value. Base64-encoded symmetric key generated by Microsoft Graph to encrypt the data value and to generate the data signature. This key is encrypted with the certificate public key that was provided during the subscription. It must be decrypted with the certificate private key before it can be used to decrypt the data or verify the signature. This key has been encrypted with the following cipher suite: RSA/ECB/OAEPWithSHA1AndMGF1Padding.
func (m *ChangeNotificationEncryptedContent) GetDataKey()(*string) {
    return m.dataKey
}
// GetDataSignature gets the dataSignature property value. Base64-encoded HMAC-SHA256 hash of the data for validation purposes.
func (m *ChangeNotificationEncryptedContent) GetDataSignature()(*string) {
    return m.dataSignature
}
// GetEncryptionCertificateId gets the encryptionCertificateId property value. ID of the certificate used to encrypt the dataKey.
func (m *ChangeNotificationEncryptedContent) GetEncryptionCertificateId()(*string) {
    return m.encryptionCertificateId
}
// GetEncryptionCertificateThumbprint gets the encryptionCertificateThumbprint property value. Hexadecimal representation of the thumbprint of the certificate used to encrypt the dataKey.
func (m *ChangeNotificationEncryptedContent) GetEncryptionCertificateThumbprint()(*string) {
    return m.encryptionCertificateThumbprint
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeNotificationEncryptedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
        }
        return nil
    }
    res["dataKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataKey(val)
        }
        return nil
    }
    res["dataSignature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataSignature(val)
        }
        return nil
    }
    res["encryptionCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificateId(val)
        }
        return nil
    }
    res["encryptionCertificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionCertificateThumbprint(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChangeNotificationEncryptedContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ChangeNotificationEncryptedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataKey", m.GetDataKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataSignature", m.GetDataSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encryptionCertificateId", m.GetEncryptionCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("encryptionCertificateThumbprint", m.GetEncryptionCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeNotificationEncryptedContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. Base64-encoded encrypted data that produces a full resource respresented as JSON. The data has been encrypted with the provided dataKey using an AES/CBC/PKCS5PADDING cipher suite.
func (m *ChangeNotificationEncryptedContent) SetData(value *string)() {
    m.data = value
}
// SetDataKey sets the dataKey property value. Base64-encoded symmetric key generated by Microsoft Graph to encrypt the data value and to generate the data signature. This key is encrypted with the certificate public key that was provided during the subscription. It must be decrypted with the certificate private key before it can be used to decrypt the data or verify the signature. This key has been encrypted with the following cipher suite: RSA/ECB/OAEPWithSHA1AndMGF1Padding.
func (m *ChangeNotificationEncryptedContent) SetDataKey(value *string)() {
    m.dataKey = value
}
// SetDataSignature sets the dataSignature property value. Base64-encoded HMAC-SHA256 hash of the data for validation purposes.
func (m *ChangeNotificationEncryptedContent) SetDataSignature(value *string)() {
    m.dataSignature = value
}
// SetEncryptionCertificateId sets the encryptionCertificateId property value. ID of the certificate used to encrypt the dataKey.
func (m *ChangeNotificationEncryptedContent) SetEncryptionCertificateId(value *string)() {
    m.encryptionCertificateId = value
}
// SetEncryptionCertificateThumbprint sets the encryptionCertificateThumbprint property value. Hexadecimal representation of the thumbprint of the certificate used to encrypt the dataKey.
func (m *ChangeNotificationEncryptedContent) SetEncryptionCertificateThumbprint(value *string)() {
    m.encryptionCertificateThumbprint = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChangeNotificationEncryptedContent) SetOdataType(value *string)() {
    m.odataType = value
}
