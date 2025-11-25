export const Framed = ({
    image,
    background,
    alt = "Framed content",
    width = "100%",
    borderWidth = "10px",
    outerRadius = "12px",
    innerRadius = "10px",
    link,
}) => {
    const imageElement = (
        <img
            src={image}
            alt={alt}
            noZoom={link ? true : false}
            style={{
                borderRadius: innerRadius,
                margin: `0 0`,
                width: "100%",
                display: "block",
            }}
        />
    );

    return (
        <div
            style={{
                padding: `${borderWidth}`,
                backgroundImage: background ? `url(${background})` : undefined,
                backgroundSize: "cover",
                backgroundPosition: "center",
                borderRadius: outerRadius,
                width: width,
            }}
        >
            <div className="overflow-hidden" >
                {link ? (
                    <a href={link}> {imageElement} </a>
                ) : (
                    imageElement
                )}
            </div>
        </div>
    );
};
